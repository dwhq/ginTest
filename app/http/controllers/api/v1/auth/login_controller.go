package auth

import (
	"gintest/app/http/controllers/api"
	"gintest/app/requests"
	services_auth "gintest/app/services/auth"
	"gintest/pkg/apiResponse"
	"gintest/pkg/auth"
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

	data, err := services_auth.Login(request.Account, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, err.Error())
	} else {
		apiResponse.Data(c, data)
	}
}

// Me 会员信息
func (lc *LoginController) Me(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// GenerateAccount 生成会员编号
func (lc *LoginController) GenerateAccount(c *gin.Context) {
	userModel := services_auth.GenerateAccount()
	response.Data(c, userModel)
}

// Register  注册会员
func (lc *LoginController) Register(c *gin.Context) {
	request := requests.RegisterRequest{}
	// 1. 验证表单
	if ok := requests.Validate(c, &request, requests.Register); !ok {
		return
	}
	err := services_auth.Register(request)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, err.Error())
	} else {
		apiResponse.Created(c, map[string]any{})
	}
}
