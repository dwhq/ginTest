package middlewares

import (
	"fmt"
	"gintest/app/models/user"
	"gintest/pkg/database"
	"gintest/pkg/jwt"
	"gintest/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			fmt.Print("cuouwxinxi", err)
			response.Unauthorized(c, err.Error())
			return
		}

		// JWT 解析成功，设置用户信息
		var userModel user.User
		database.DB.Where("id", claims.UserID).First(&userModel)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_account", userModel.Account)
		c.Set("current_user", userModel)

		c.Next()
	}
}
