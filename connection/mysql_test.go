package connection

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"Username" form:"Username"`
	Password string `json:"password" form:"password"`
}

func (u User) TableName() string {
	return "user"
}

func getNewDB() (D *gorm.DB) {
	return GetMysql()
}
func TestMysql(t *testing.T) {
	var user User
	db := getNewDB()
	result := db.Table("user").First(&user, "id=?", "1")
	if result.Error != nil {
		fmt.Println("查询错误")
	}
	if result.RowsAffected == 0 { //返回受影响的行数
		fmt.Println("没有匹配的记录")
	} else {
		fmt.Println("查询到记录", user)
	}
}
func TestMysql2(t *testing.T) {
	var user User
	db := getNewDB()
	db.Find(&user, "3", user.TableName())
	fmt.Println("数据：", user)
}
