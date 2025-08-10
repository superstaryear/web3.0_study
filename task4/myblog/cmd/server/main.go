package main

import (
	"github.com/gin-gonic/gin"
	"myblog/api"
	"myblog/middleware"
)

func register(c *gin.Context) {
	api.Register(c)
}

func login(c *gin.Context) {
	api.Login(c)
}

func main() {
	router := gin.Default()
	//外部不需要token的分组成outside前缀,引入错误和日志中间件
	outside := router.Group("outside", middleware.ErrorHandler, middleware.LogHandler())
	//用户注册
	outside.POST("/register", register)
	//用户登录
	outside.POST("/login", login)
	router.Run(":8080")
}
