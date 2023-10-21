package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
