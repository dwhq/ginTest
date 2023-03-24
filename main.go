package main

import (
	"gintest/bootstrap"
	"gintest/config"
	c "gintest/pkg/config"
	"gintest/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	// 初始化 Logger
	bootstrap.SetupLogger()
}

func main() {
	bootstrap.SetupDB()
	r := gin.Default()
	routes.RegisterWebRoutes(r)//路由
	r.Static("/assets", "./assets")
	if err := r.Run(":"+c.GetString("app.port")); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}