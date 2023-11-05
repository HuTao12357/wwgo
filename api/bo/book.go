package bo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"wwgo/connection"
	"wwgo/global"
)

type Book struct {
	Id       int    `json:"id" form:"id"`
	bookName string `json:"book_name" form:"book_name"`
}

func getNewDB() (D *gorm.DB) {
	return connection.GetMysql()
}
func GetById(c *gin.Context) {
	id := c.Query("id")
	var book Book
	//sql := fmt.Sprint
	global.GLOBAL_DB.Table("book").Where(id, "1").Find(&book)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": book,
	})
}
