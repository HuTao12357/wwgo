package test

import (
	"fmt"
	"testing"
	"wwgo/utils"
)

func Test_token(t *testing.T) {
	username := "张三"
	id := "666666"
	res, err := utils.GenToken(username, id)
	if err != nil {
		fmt.Println("生成token失败")
	}
	fmt.Println("生成token：", res)
}
