package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	// "crypto/rsa"  
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"encoding/base64"
)

// 参数i
func rand_char(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 填充 16位字符
// func PKCS5Padding(plaintext []byte, blockSize int) []byte {
// 	padding := blockSize - len(plaintext)%blockSize//padding的值就是16 - x
// 	padText := bytes.Repeat([]byte{byte(padding)}, padding)//要填充的字节:padding个值为padding的字节
// 	return append(plaintext, padText...)//填充在明文后面
// }
func PaddingText1(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	//fmt.Println(newPaddingStr)
	return newPaddingStr
}

// aes cbc 加密
// 加密算法解析
// func AesEncrypt(key, plaintext []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
// 	ciphertext := make([]byte, len(plaintext)+aes.BlockSize)
// 	// iv ="0102030405060708"
// 	iv := ciphertext[:aes.BlockSize]
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		panic(err)
// 	}
// 	mode := cipher.NewCBCEncrypter(block, iv)
// 	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext[:])
// 	return ciphertext, nil
// }

func AesEncrypt(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	src = PaddingText1(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte("0102030405060708"))
	blockMode.CryptBlocks(src, src)
	return src
}

// 生成params
func gen_params(d []byte, g string, i string) []byte {
	fmt.Println(string(d))
	t := AesEncrypt(d, []byte(g))
	text := AesEncrypt(t, []byte(i))
	return text
}

/*
2次AES加密中，初始向量都是0102030405060708，加密模式都是CBC加密，
不同的是第一次加密中，d作为message，g作为key来加密；第二次加密中，把第一次加密结果作为message，
i作为key来加密。我们可以通过Crypto.Cipher中的AES实现，

d   "{"rid":"R_SO_4_1454730043","offset":"20","total":"false","limit":"20","csrf_token":""}"
g   "0CoJUm6Qyw8W8jud"

*/
// func rsa_encrypt(msg string) string{
// 	cryptor := rsa.
// }
func main() {
	i := rand_char(16)
	g := "0CoJUm6Qyw8W8jud"
	query := map[string]string{
		"rid":        "R_SO_4_1454730043",
		"offset":     "20",
		"total":      "false",
		"limit":      "20",
		"csrf_token": "",
	}

	query_json, _ := json.Marshal(query)
	// query_string := string(query_json)
	// query_byte := []byte(query_string)

	fmt.Println(i)
	// fmt.Println("query_string:",query_string)
	// fmt.Println(query_byte)
	p := gen_params(query_json, g, i)
	params := base64.StdEncoding.EncodeToString(p)
	fmt.Println(params)
}
