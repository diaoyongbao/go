package main

// BUG 解决token无法生成的问题
// TODO 生成测试报告
import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

func NewToken() string {
	key := "秘钥"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 自定义键值对
		"foo": "xxx",
		"nbf": time.Now().Add(TokenExpireDuration).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(key)) // 传入秘钥 []byte() 拿到string
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func decode(tokenstring string) {

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("秘钥"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 获取分区内容
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}

func TestNewToken(t *testing.T) {
	tokenstring := NewToken()
	t.Log(tokenstring)
}
