package config

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"Username"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
}

func (u *User) TableName() string {
	return "user"
}
func TestMysqlGet(t *testing.T) {
	var user User
	db, err := MysqlGet()
	if err != nil {
		log.Fatal(err)
	}
	db.Where("id = ?", "3").First(&user)
	fmt.Println("数据：", user)
}

func TestRedisGet(t *testing.T) {
	rdb := GetRedis()
	rdb.Set("money", 6999, 0)
	v := rdb.Get("money")
	fmt.Println(v)
}
func TestYmlPath(t *testing.T) {
	var s = "./config.yml"
	//得到绝对路径
	s1, _ := filepath.Abs(s)
	fmt.Println(s1)
}
