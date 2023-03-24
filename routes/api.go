package routes

import (
	v1 "gintest/app/http/controllers/api/v1"
	"gintest/app/http/controllers/api/v1/aouth"
	shop "gintest/app/http/controllers/api/v1/shop"
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(router *gin.Engine) {
	page := new(v1.ArticlesController)
	router.GET("/index", page.Index)//系统首页
	product := new(shop.ProductController)
	router.GET("/product", product.Index)//商品列表
	//用户认证
	login := new(aouth.LoginController)
	router.POST("/auth/login", login.Login)//商品列表


}