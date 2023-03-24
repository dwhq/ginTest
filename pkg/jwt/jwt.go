// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"gintest/pkg/app"
	"gintest/pkg/config"
	"gintest/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// JWT 定义一个jwt对象
type JWT struct {

	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

// MyCustomClaims 自定义载荷
type MyCustomClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

// ParserToken 解析 Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*MyCustomClaims, error) {

	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 2. 解析出错
	if err != nil {
		//validationErr, ok := err.()
		//if ok {
		//	if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
		//		return nil, ErrTokenMalformed
		//	} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
		//		return nil, ErrTokenExpired
		//	}
		//}
		if errors.Is(err, jwtpkg.ErrTokenMalformed) {
			return nil, ErrTokenInvalid
		} else if errors.Is(err, jwtpkg.ErrTokenExpired) || errors.Is(err, jwtpkg.ErrTokenNotValidYet) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {

	// 1. 从 Header 里获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		return "", err
		//validationErr, ok := err.(*jwtpkg.ValidationError)
		//// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		//if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
		//	return "", err
		//}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*MyCustomClaims)

	// 5. 检查是否过了『最大允许刷新的时间』
	x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Unix() > x {
		// 修改过期时间
		claims.RegisteredClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 生成  Token，在登录成功时调用
func (jwt *JWT) IssueToken(userID string, userName string) (token string, expireAtTime *jwtpkg.NumericDate) {

	// 1. 构造用户 claims 信息(负荷)
	expireAtTime = jwt.expireAtTime()
	claims := MyCustomClaims{
		userID,
		userName,
		jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimenowInTimezone()), // 签名生效时间
			IssuedAt:  jwtpkg.NewNumericDate(app.TimenowInTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireAtTime,                                   // 签名过期时间
			Issuer:    config.GetString("app.name"),                   // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return
	}

	return
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims MyCustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() *jwtpkg.NumericDate {
	timeNow := app.TimenowInTimezone()
	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute
	return jwtpkg.NewNumericDate(timeNow.Add(expire))
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// getTokenFromHeader 使用 jwtpkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
