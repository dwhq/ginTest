package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RegisterRequest struct {
	Account    string `json:"account,omitempty"  valid:"account"`
	Mobile     string `json:"mobile,omitempty" valid:"mobile"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
	PayPass    string `json:"pay_pass,omitempty" valid:"pay_pass"`
	InviteCode string `json:"invite_code,omitempty" valid:"invite_code"`
}

func Register(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"account":     []string{"required", "min:3"},
		"mobile":      []string{"required", "digits:11"},
		"verify_code": []string{"required"},
		"password":    []string{"required"},
		"pay_pass":    []string{"required"},
		"invite_code": []string{"required"},
	}
	messages := govalidator.MapData{
		"account": []string{
			"required:会员编号不能为空",
			"min:用户名长度需大于 3",
		},
		"mobile": []string{
			"required:手机号不能为空",
			"min:用户名长度需大于 3",
		},
		"verify_code": []string{
			"required:邀请码不能为空",
		},
		"password": []string{
			"required:密码不能为空",
		},
		"pay_pass": []string{
			"required:安全密码不能为空",
		},
		"invite_code": []string{
			"required:邀请码不能为空",
		},
	}
	return validate(data, rules, messages)
}
