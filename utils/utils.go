package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 使用bcrypt加密密码,每一次都是不同的结果
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Sprintf("发生错误")
	}
	return string(hash)
}

// 验证密码,正确返回true，错误返回false
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// GenToken 生成JWT
const TokenExpireDuration = time.Hour * 2 //过期时间
var mySecret = []byte("123456")           //密钥
type MyClaims struct {
	Username string `json:"username"`
	UserId   string `json:"userid"`
	jwt.StandardClaims
}

func GenToken(username string, userid string) (string, error) {
	fmt.Println("===================a", TokenExpireDuration)
	c := MyClaims{
		username,
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "myProject",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tmp, err := token.SignedString(mySecret)
	return tmp, err
}

// 将字符进行大小写转换
func StringCase(str string) (restr string) {
	char := []rune(str) //rune int32，一般用来表示unicode
	for k, v := range char {
		if v >= 65 && v <= 90 {
			char[k] = v + 32 //26个英文字母
		} else if v >= 97 && v <= 122 {
			char[k] = v - 32
		} else {
			return "字符串超出范围"
		}
	}
	return string(char)
}
