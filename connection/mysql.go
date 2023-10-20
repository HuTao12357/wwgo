package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var db *gorm.DB

func GetMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/wwgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 处理连接错误
		fmt.Println("连接错误==================")
	}

}
func init() {
	GetMysql()
}
