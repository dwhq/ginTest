package v1

import (
	"gintest/app/http/controllers/api"
	"gintest/pkg/apiResponse"
	"github.com/gin-gonic/gin"
	"gintest/app/services/page"

)
// ArticlesController 处理静态页面
type ArticlesController struct {
	api.BaseController
}

// Index 系统首页
func (ac *ArticlesController) Index(c *gin.Context) {
	apiResponse.Data(c, page.Index())
}