package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserInfo struct {
	Username string `json:"user"`
	Password string `json:"passwd"`
}

// 定义token过期时间
const TokenExpireDuration = time.Hour * 2

//Mysecret  密钥
const (
	SecretKey = "我的密钥"
)

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个自己的声明
	fmt.Println(username)
	c := MyClaims{
		"username",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "my-project",
		},
	}
	fmt.Println(c)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// fmt.Println(token)
	fmt.Println(token.SignedString([]byte(SecretKey)))
	return token.SignedString([]byte(SecretKey))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}

// 鉴权接口
func authHandler(c *gin.Context) {
	var user UserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	// 检验用户名及密码是否正确
	if user.Username == "dyb" && user.Password == "123456" {
		tokenString, _ := GenToken(user.Username)
		// fmt.Println(user.Username, user.Password 
		// fmt.Println(tokenString)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

// token校验中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式，1. 放在请求头 2， 放在请求体 3，放在URI
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按照空分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，使用之前定义的jwt进行解析
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"mgs":  "无效的token",
			})
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}

func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}

func main() {
	r := gin.Default()
	// fmt.Println(("dyb"))
	r.POST("/auth", authHandler)
	r.GET("/home", JWTAuthMiddleware(), homeHandler)
	r.Run()
}
