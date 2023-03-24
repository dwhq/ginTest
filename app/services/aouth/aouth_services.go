package aouth

import (
	"errors"
	"gintest/app/models/user"
	"gintest/pkg/database"
	"gintest/pkg/hash"
	"gintest/pkg/jwt"
	"math/rand"
)

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
func GenerateAccount() (data map[string]any) {
	var account int64
	for i := 0; i < 10; i++ {
		account = rand.Int63n(10000000)
		var count int64
		database.DB.Where("account=?", account).Count(&count)
		if count == 0 {
			break
		}
	}
	data = map[string]any{
		"account": account,
	}
	return
}
