package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"myblog/common"
	"time"
)

/*
日志记录中间件
*/
func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 记录日志
		common.LogFile.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),    // 状态码
			"method":  c.Request.Method,     // 请求方法
			"path":    c.Request.URL.Path,   // 请求路径
			"latency": latencyTime.String(), // 延迟时间
		}).Info("HTTP请求")
	}
}
