package shop

import (
	"gintest/app/http/controllers/api"
	"gintest/app/services/shop"
	"gintest/pkg/apiResponse"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductController struct {
	api.BaseController
}

func (p *ProductController) Index(c *gin.Context)  {
	categoryId := c.DefaultQuery("category_id","0")
	curType := c.DefaultQuery("cur_type","1")
	intId, _:=strconv.Atoi(categoryId)
	intCurType, _:=strconv.Atoi(curType)
	apiResponse.Data(c,shop.Index(c,10,intId,intCurType))
}