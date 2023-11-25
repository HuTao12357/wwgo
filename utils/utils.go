package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

// HashAndSalt 使用bcrypt加密密码,每一次都是不同的结果
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("加密过程发生错误")
	}
	return string(hash)
}

// ComparePasswords 验证密码,正确返回true，错误返回false
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// TokenExpireDuration GenToken 生成JWT
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

// StringCase 将字符进行大小写转换
func StringCase(str string) (res string) {
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

// IsEmptyString go没有工具类直接判断是否为空
func IsEmptyString(str string) bool {
	if len(str) == 0 {
		return false
	} else {
		return true
	}
}

// IsEmptyArray 切片
func IsEmptyArray(str interface{}) bool {
	//只能使用断言了，go还不支持原生泛型
	switch arr := str.(type) {
	case []int:
		return len(arr) == 0
	case []string:
		return len(arr) == 0
	case []float64:
		return len(arr) == 0
	case []float32:
		return len(arr) == 0
	case []bool:
		return len(arr) == 0
	}
	return false
}

// StringAdd 字符串相加
func StringAdd(s, str string) string {
	var result strings.Builder
	result.WriteString(s)
	result.WriteString(str)
	s2 := result.String()
	return s2
}

// CreatFile 创建文件
func CreatFile(fileName string) string {
	file, err := os.Create(fileName)
	if err != nil {
		return "创建文件失败"
	}
	defer file.Close()
	str := fmt.Sprintf("创建文件：%s成功", fileName)
	return str
}

// WriteContent 清空文件，并写入内容
func WriteContent(content, fileName string) string {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	//os.O_TRUNC：清除文件以前内容
	if err != nil {
		return "打开文件失败"
	}
	_, err = file.WriteString(content)
	file.WriteString("\n")
	if err != nil {
		return "写入失败"
	}
	defer file.Close()
	return "写入成功"
}

// WriteContentAdd 追加内容
func WriteContentAdd(content, fileName string) string {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	/**
	os.O_WRONLY：表示只写方式打开文件
	os.O_CREATE：表示没有就创建
	os.O_APPEND：在文件末尾添加文件
	*/
	if err != nil {
		return "打开文件失败"
	}
	_, err = file.WriteString(content)
	file.WriteString("\n")
	if err != nil {
		return "写入失败"
	}
	defer file.Close()
	return "写入成功"
}

// ReadFile 读取文件
func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "读取文件失败"
	}
	return string(file)
}

// DeleteFile 删除文件
func DeleteFile(fileName string) bool {
	err := os.Remove(fileName)
	if err != nil {
		return false
	}
	return true
}

// GetTimeNowHM 获取当前时间string，带有时分秒
func GetTimeNowHM() string {
	ft := fmt.Sprintf("2006-01-02 15:04:05")
	timeNow := time.Now().Format(ft)
	return timeNow
}

// GetTimeNow 获取当前时间string，没有时分秒
func GetTimeNow() string {
	ft := fmt.Sprintf("2006-01-02")
	timeNow := time.Now().Format(ft)
	return timeNow
}

// GetCurrentTimestamp 获取当前时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}
