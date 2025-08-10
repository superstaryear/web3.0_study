package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"myblog/common"
)

/*
全局异常处理中间件 调用需要放在所有中间件第一的位置
*/
func ErrorHandler(c *gin.Context) {
	// 先执行后续中间件
	c.Next()
	if len(c.Errors) > 0 {
		err := c.Errors.Last()
		var message string
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			message = "资源不存在"
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			message = "用户名或密码错误"
		default:
			message = "服务器内部错误:" + err.Error()
		}
		//记录日志
		common.LogFile.WithFields(logrus.Fields{
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		}).Error(message)
		common.Error(c, 500, message)
	}
}
