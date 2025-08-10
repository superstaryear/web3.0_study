package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"myblog/common"
	"myblog/middleware"
	"myblog/model"
	"time"
)

/*
用户注册
*/
func Register(c *gin.Context) {
	var registerReq model.RegisterRequest
	if err := c.ShouldBind(&registerReq); err != nil {
		errMsg := common.GetValidMsg(err, &registerReq)
		common.Error(c, 500, "注册请求参数错误: "+errMsg)
		return
	}
	var userEntity model.UserEntity
	copier.Copy(&userEntity, &registerReq)
	fmt.Println(userEntity)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	userEntity.Password = string(hashedPassword)
	if err := common.DB.Create(&userEntity).Error; err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	var registerRes model.RegsiterResponse = model.RegsiterResponse{
		UserId:   userEntity.ID,
		Username: userEntity.Username,
	}
	common.Success(c, registerRes, "注册成功")
	return
}

/*
用户登录
*/
func Login(c *gin.Context) {
	var loginReq model.LoginRequest
	if err := c.ShouldBind(&loginReq); err != nil {
		errMsg := common.GetValidMsg(err, &loginReq)
		common.Error(c, 500, errMsg)
		return
	}
	//查询用户
	var userEntity model.UserEntity
	err := common.DB.Where("username = ?", loginReq.Username).Take(&userEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Error(c, 500, "用户名或密码错误")
			return
		}
		common.Error(c, 500, fmt.Sprintf("查询失败:%s", err.Error()))
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(loginReq.Password)); err != nil {
		common.Error(c, 500, "用户名或密码错误")
		return
	}
	tokenStr, err := middleware.GenerateToken(userEntity.ID, userEntity.Username, time.Now().Add(8*time.Hour).Unix())
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	var loginRes model.LoginResponse = model.LoginResponse{
		Token: tokenStr,
	}
	common.Success(c, loginRes, "登录成功")
}
