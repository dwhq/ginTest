package routes

import (
	v1 "gintest/app/http/controllers/api/v1"
	"gintest/app/http/controllers/api/v1/auth"
	shop "gintest/app/http/controllers/api/v1/shop"
	"gintest/app/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(router *gin.Engine) {
	page := new(v1.ArticlesController)
	router.GET("/index", page.Index) //系统首页
	product := new(shop.ProductController)
	router.GET("/product", product.Index) //商品列表
	//用户认证
	login := new(auth.LoginController)
	router.POST("/auth/login", login.Login)                     //登录
	router.POST("/auth/register", login.Register)               //注册
	router.GET("/auth/generate_account", login.GenerateAccount) //生成会员编号

	router.GET("/auth/me", middlewares.AuthJWT(), login.Me) //用户信息

}
