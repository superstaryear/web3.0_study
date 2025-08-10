package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

/*
*
全局统一结构体返回格式
*/
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(200, Response{Code: 200, Data: data, Message: message})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(200, Response{Code: code, Message: msg})
}

// 校验框架返回中文提示
func GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}
