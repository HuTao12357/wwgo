package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var err error
var db *gorm.DB

var dbName = os.Getenv("DB_NAME")
var dbPassword = os.Getenv("DB_PASSWORD")

const port = "3306"

func GetMysql() (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/wwgo?charset=utf8mb4&parseTime=True&loc=Local", dbName, dbPassword, port)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 处理连接错误
		fmt.Println("连接错误==================")
	}
	return
}
func init() {
	GetMysql()
}
