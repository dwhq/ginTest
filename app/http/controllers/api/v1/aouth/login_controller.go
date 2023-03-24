package aouth

import (
	"gintest/app/http/controllers/api"
	"gintest/app/requests"
	"gintest/app/services/aouth"
	"gintest/pkg/apiResponse"
	"gintest/pkg/response"
	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	api.BaseController
}

// Login 会员登录

func (lc *LoginController) Login(c *gin.Context) {
	request := requests.LoginRequest{}
	// 1. 验证表单
	if ok := requests.Validate(c, &request, requests.Login); !ok {
		return
	}

	data,err := aouth.Login(request.Account,request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, err.Error())
	}else {
		apiResponse.Data(c,data)
	}
}
