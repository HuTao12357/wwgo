package bo

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math"
	"net/http"
	"wwgo/global"
)

type Book struct {
	Id       int             `json:"id" form:"id"`
	BookName string          `json:"book_name" form:"book_name"`
	Money    decimal.Decimal `json:"money" form:"money"`
	Num      int             `json:"num" form:"num"`
	Rate     float64
}

func GetById(c *gin.Context) {
	id := c.Query("id")
	var book Book
	//sql := fmt.Sprint
	global.GlobalDB.Table("book").Where(id, "1").Find(&book)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": book,
	})
}
func GetNumRate(c *gin.Context) {
	var book []Book
	result := global.GlobalDB.Table("book").Find(&book)
	if result.RowsAffected == 0 {
		panic("表中没有数据")
	}

	var sum int
	if err := global.GlobalDB.Table("book").Select("sum(num)").Scan(&sum).Error; err != nil {
		panic("failed to query database")
	}

	for k := range book {
		rate := float64(book[k].Num) / float64(sum) * 100 //百分比
		book[k].Rate = math.Round(rate*100) / 100         //保留两位小数
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": book, //应该返回book切片，而不是result
	})
}
