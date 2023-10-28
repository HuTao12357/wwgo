package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
	"wwgo/utils"
)

func TestBcrypt(t *testing.T) {
	//测试Bcrypt
	word := "123456"
	JMWord := utils.HashAndSalt([]byte(word))
	fmt.Println("加密后密码：", JMWord)
	a := utils.ComparePasswords(JMWord, []byte(word))
	fmt.Println("a==", a)
	//测试uuid
	ids := uuid.NewV4().String()
	fmt.Println("生成的uuid:", ids)
}
