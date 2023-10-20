package us

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"wwgo/connection"
)

func getNewDB() (D *gorm.DB) {
	return connection.GetMysql()
}

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"Username"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "无效参数",
		})
		return
	}
	name := user.Username
	password := user.Password
	if name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号或密码不能为空",
		})
		return
	}
	db := getNewDB()
	//sql := fmt.Sprintf("select * from admin where username='%s' ", user.Username) //原生sql
	db.Where("username = ?", name).First(&user)
	fmt.Println("========", user)
}
