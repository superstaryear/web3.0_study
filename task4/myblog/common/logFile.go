package common

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var LogFile = logrus.New()

/*
初始化日志文件
*/
func init() {
	// 确保日志目录存在
	if err := os.MkdirAll("logs", 0755); err != nil {
		logrus.Fatal("无法创建日志目录: ", err)
	}
	// 检查文件是否可写
	_, err := os.OpenFile("logs/blog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("无法打开日志文件: ", err)
	}
	// 输出到文件
	LogFile.SetOutput(&lumberjack.Logger{
		Filename:   "logs/blog.log", // 日志文件路径
		MaxSize:    10,              // 单个日志文件最大大小(MB)
		MaxBackups: 5,               // 保留旧日志文件的最大数量
		MaxAge:     30,              // 保留旧日志文件的最大天数
		Compress:   true,            // 是否压缩/归档旧日志
	})

	// 设置文本格式
	LogFile.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 时间格式
		FullTimestamp:   true,                  // 显示完整时间
		DisableColors:   false,                 // 文件输出时自动禁用颜色
	})

	// 设置日志级别
	LogFile.SetLevel(logrus.InfoLevel)
	//LogFile.SetLevel(logrus.ErrorLevel)
}
