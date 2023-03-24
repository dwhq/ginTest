package aouth

import (
	"errors"
	"fmt"
	"gintest/app/models/user"
	"gintest/pkg/database"
	"gintest/pkg/hash"
	"gintest/pkg/jwt"
)

func Login(account string,password string) (data map[string]any,err error) {
	var userInfo  user.User
	database.DB.Where("account=?",account).First(&userInfo)
	if userInfo.ID == 0 {
		return data ,errors.New("账号不存在")
	}
	if !hash.BcryptCheck(password, userInfo.Password) {
		return data, errors.New("密码错误")
	}
	token,expireAtTime := jwt.NewJWT().IssueToken(string(userInfo.ID), userInfo.Account)
	fmt.Print(jwt.MyCustomClaims{});
	data = map[string]any{
		"token" :token,
		"token_type" : "Bearer",
		"expires_in": expireAtTime,
	}
	return
}
