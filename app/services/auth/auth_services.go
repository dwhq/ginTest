package auth

import (
	"errors"
	"fmt"
	"gintest/app/models/user"
	"gintest/app/requests"
	"gintest/pkg/database"
	"gintest/pkg/hash"
	"gintest/pkg/jwt"
	"math/rand"
)

// Login 登录
func Login(account string, password string) (data map[string]any, err error) {
	var userInfo user.User
	database.DB.Where("account=?", account).First(&userInfo)
	if userInfo.ID == 0 {
		return data, errors.New("账号不存在")
	}
	if !hash.BcryptCheck(password, userInfo.Password) {
		return data, errors.New("密码错误")
	}
	token, expireAtTime := jwt.NewJWT().IssueToken(userInfo.GetStringID(), userInfo.Account)
	data = map[string]any{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   expireAtTime,
	}
	return
}

// Register 注册
func Register(requests requests.RegisterRequest) (err error) {
	var (
		userInfo user.User
		count    int64
		reInfo   user.User
	)

	database.DB.Where("account=?", requests.InviteCode).First(&reInfo)
	if reInfo.ID == 0 {
		return errors.New("推荐人不存在")
	}
	database.DB.Model(&user.User{}).Where("account=?", requests.Account).Count(&count)
	fmt.Print("总数", count, requests.Account)
	if count > 0 {
		return errors.New("账号已存在")
	}
	rePath := ""
	if len(reInfo.RePath) > 0 {
		rePath = reInfo.RePath + "," + reInfo.GetStringID()
	} else {
		rePath = reInfo.GetStringID()
	}
	userInfo.Account = requests.Account
	userInfo.Mobile = requests.Mobile
	userInfo.Nickname = "GY" + requests.Account
	userInfo.Gender = 0
	userInfo.Password = hash.BcryptHash(requests.Password)
	userInfo.PayPass = hash.BcryptHash(requests.PayPass)
	userInfo.AvatarUrl = "images/user_head.png"
	userInfo.ReId = reInfo.ReId
	userInfo.ReName = reInfo.Account
	userInfo.ReLevel = reInfo.ReLevel + 1
	userInfo.RePath = rePath
	userInfo.IsLock = 0
	//userInfo.RunRegisterAt = null
	database.DB.Create(&userInfo)
	if userInfo.ID == 0 {
		return errors.New("创建用户失败，请稍后尝试~")
	}
	return
}

// GenerateAccount  生成会员编号
func GenerateAccount() (data map[string]any) {
	var account int64
	for i := 0; i < 10; i++ {
		account = rand.Int63n(10000000)
		var count int64
		database.DB.Model(&user.User{}).Where("account=?", account).Count(&count)
		if count == 0 {
			break
		}
	}
	data = map[string]any{
		"account": account,
	}
	return
}
