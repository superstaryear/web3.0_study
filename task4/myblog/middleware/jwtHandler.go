package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"myblog/common"
	"strings"
)

var privateKey []byte = []byte("private key")

/*
生成jwt Token
*/
func GenerateToken(userId uint, userName string, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer":   "blog", //发行者
		"userName": userName,
		"userId":   userId,
		"exp":      exp, // 过期时间
	})
	tokenStr, err := token.SignedString(privateKey)
	return tokenStr, err
}

/*
*
jwt 鉴权中间件
*/
func JwtHandler(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		common.Error(c, 500, "token不存在")
		c.Abort()
		return
	}
	tokenStr := strings.Split(authorization, " ")[1]
	if tokenStr == "" {
		common.Error(c, 500, "token非法")
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		//确保签名方法是我们期望的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	if err != nil {
		common.Error(c, 500, "token非法")
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		issuer := claims["issuer"]
		//判断签名主体是否一致
		if issuer != "blog" {
			common.Error(c, 500, "token非法")
			c.Abort()
			return
		}
		c.Set("userId", claims["userId"])
		c.Set("userName", claims["userName"])
		c.Next()
	} else {
		common.Error(c, 500, "token不存在")
		c.Abort()
		return
	}
}
