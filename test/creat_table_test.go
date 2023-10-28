package test

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"testing"
	"wwgo/connection"
)

type Book struct {
	Id     string
	Name   string
	Author string
	Detail string
}

func Test_GormTable(t *testing.T) {

	db := getNewDB()
	M := db.Migrator()
	M.CreateTable(&Book{})
}
func Test_DropTable(t *testing.T) {
	db := getNewDB()
	M := db.Migrator()
	//M.DropTable(&Book{}) 		//删除表
	err := M.DropColumn(&Book{}, "price") //删除列
	if err != nil {
		fmt.Println("执行出错：", err)
	}
}
func Test_Column(t *testing.T) {
	db := getNewDB()
	//db.Migrator().AddColumn(&Book{}, "detail")
	id, _ := uuid.NewRandom()
	ids := id.String()
	book := Book{Id: ids, Name: "哈姆雷特", Author: "莎士比亚", Detail: "复仇"}
	db.Table("gorm_book").Create(&book)
}
func getNewDB() (D *gorm.DB) {
	return connection.GetMysql()
}
