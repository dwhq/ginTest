package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginRequest struct {
	Account  string `json:"account,omitempty"  valid:"account"`
	Password string `json:"password,omitempty" valid:"password"`
}

func Login(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"account":  []string{"required"},
		"password": []string{"required"},
	}
	messages := govalidator.MapData{
		"account": []string{
			"account:用户名为必填项",
			"min:用户名长度需大于 3",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}
	return validate(data, rules, messages)
}
