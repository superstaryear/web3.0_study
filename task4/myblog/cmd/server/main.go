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

func postCreate(c *gin.Context) {
	api.PostCreate(c)
}

func postQuery(c *gin.Context) {
	api.PostQuery(c)
}

func postPage(c *gin.Context) {
	api.PostPage(c)
}

func postUpdate(c *gin.Context) {
	api.PostUpdate(c)
}

func postDelete(c *gin.Context) {
	api.PostDelete(c)
}

func commentCreate(c *gin.Context) {
	api.CommentCreate(c)
}

func commentQuery(c *gin.Context) {
	api.CommentQuery(c)
}

func main() {
	router := gin.Default()
	//外部不需要token的分组成outside前缀,引入错误和日志中间件
	outside := router.Group("outside", middleware.ErrorHandler, middleware.LogHandler())
	{
		//用户注册
		outside.POST("/register", register)
		//用户登录
		outside.POST("/login", login)
	}
	//api接口需要jwt鉴权校验分组成api前缀，引入错误、日志、jwt鉴权校验中间件
	api := router.Group("api", middleware.ErrorHandler, middleware.LogHandler(), middleware.JwtHandler)
	{
		//文章创建
		api.POST("/post/create", postCreate)
		//查询单个文章的详细信息
		api.GET("/post/:postId", postQuery)
		//获取所有文章列表
		api.POST("/post/page", postPage)
		//文章更新
		api.POST("/post/update", postUpdate)
		//文章批量删除
		api.POST("/post/delete", postDelete)
		//评论创建
		api.POST("/comment/create", commentCreate)
		//评论列表查询
		api.POST("/comment/list", commentQuery)
	}
	//启动服务 8080端口
	router.Run(":8080")
}
