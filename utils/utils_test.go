package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestBcrypt(t *testing.T) {
	//测试Bcrypt
	word := "123456"
	JMWord := HashAndSalt([]byte(word))
	fmt.Println("加密后密码：", JMWord)
	a := ComparePasswords(JMWord, []byte(word))
	fmt.Println("a==", a)
	//测试uuid
	ids := uuid.NewV4().String()
	fmt.Println("生成的uuid:", ids)
}
