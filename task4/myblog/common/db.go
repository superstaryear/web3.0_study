package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	username = "root"      //账号
	password = "55555"     //密码
	host     = "555.5.5.5" //数据库地址，可以是Ip或者域名
	port     = 3306        //数据库端口
	Dbname   = "gorm"      //数据库名
	timeout  = "10s"       //连接超时，10秒
	DB       *gorm.DB      //全局数据库datasource
)

/*
初始化db
*/
func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	var err error
	//gorm的默认日志是只打印错误和慢SQL，这里开启sql打印
	var mysqlLogger logger.Interface
	mysqlLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	fmt.Println("Database connected!")
	//DB.AutoMigrate(&model.UserEntity{}, &model.PostEntity{}, &model.CommentEntity{})
	//fmt.Println("Database table init!")
}
