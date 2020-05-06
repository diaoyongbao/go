package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("123")

// GenToken 生成JWT
func GenToken(username string) (string,error){
	// 创建一个自己的声明

		c := MyClaims{
			"username",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
				Issuer: "my-project",
			}
		}
	token := jwt.NewWithClaims(jwt.SigningMethodES256,c)
	return token.SignedString(MySecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims,error){
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString,&Myclaims{},func(token *jwt.Token)(i interface{},err error){
		return MySecret,nil
	})
	if err != nil {
		return nil,err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid{
		return claims,nil
	}
	return nil,errors.New("invalid token")

}