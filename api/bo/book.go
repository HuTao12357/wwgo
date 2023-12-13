package bo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"log"
	"math"
	"net/http"
	"os"
	"wwgo/common"
	"wwgo/global"
)

type Book struct {
	Id       int             `json:"id" form:"id"`
	BookName string          `json:"book_name" form:"book_name"`
	Money    decimal.Decimal `json:"money" form:"money"`
	Num      int             `json:"num" form:"num"`
	Rate     float64         `gorm:"-"`
}
type Ids struct {
	Ids      []int  `json:"id" form:"id"`
	FileName string `json:"file_name" form:"file_name"`
}

func GetById(c *gin.Context) {
	id := c.Query("id")
	var book Book
	global.GlobalDB.Table("book").Where(id, "1").Find(&book)
	global.GlobalDB.Debug().Table("book").Where("id = ?", id).Find(&book) //DEBUG查看执行的sql,sql会有:表明.id,是gorm自动添加到，避免id字段存在多个表中引起的歧义
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

func GetExcel(c *gin.Context) {
	var ids Ids
	var book []Book
	if err := c.ShouldBind(&ids); err != nil {
		fmt.Println("绑定失败")
	}
	idStr := ids.Ids
	err := global.GlobalDB.Table("book").Where("id in (?)", idStr).Find(&book).Error //可以使用Omit来忽略想要显示的字段
	if err != nil {
		fmt.Println("查询失败")
	}
	sheet := "Sheet1"
	fileName := ids.FileName

	//逻辑判断
	if _, err := os.Stat(fileName); err == nil {
		log.Printf("文件名:%s,已存在", fileName)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "文件名已存在",
		})
		return
	} else if os.IsNotExist(err) {
		log.Printf("文件:%s,不存在", fileName)
	} else if os.IsPermission(err) {
		log.Println("没有权限")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "没有权限",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "无法确定",
		})
		return
	}
	//创建Excel文件
	f := excelize.NewFile()
	f.SetCellValue(sheet, "A1", "BOOK")
	f.SetCellValue(sheet, "A2", "id")
	f.SetCellValue(sheet, "B2", "name")
	f.SetCellValue(sheet, "C2", "money")
	f.SetCellValue(sheet, "D2", "num")
	//填充数据
	for k, v := range book {
		row := k + 3
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), v.Id)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), v.BookName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), v.Money)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), v.Num)
	}
	// 保存 Excel 文件
	err = f.SaveAs(fileName)
	if err != nil {
		log.Fatalf("创建Excel失败") //输出一条格式化信息，并退出程序，生产环境不要用
	} else {
		log.Println("GetExcel接口创建Excel成功")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": book,
	})
}
func TestResult(c *gin.Context) {
	data := "hello go"
	err := common.Foo()
	resp := common.Fail(data, err)
	c.JSON(http.StatusOK, resp)
}
