package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 添加这行
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	username = "root"          //账号
	password = "E05302"        //密码
	host     = "102.202.62.62" //数据库地址，可以是Ip或者域名
	port     = 3306            //数据库端口
	Dbname   = "gorm"          //数据库名
	timeout  = "10s"           //连接超时，10秒
	DB       *gorm.DB
	//DB *sqlx.DB
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	var err error
	//gorm的默认日志是只打印错误和慢SQL，这里开启sql打印
	var mysqlLogger logger.Interface
	mysqlLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	//db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	fmt.Println("Database connected!")
}
