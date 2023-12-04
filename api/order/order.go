package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math/rand"
	"net/http"
	"time"
	"wwgo/api/bo"
	"wwgo/common"
	"wwgo/global"
)

type Order struct {
	Id       int             `json:"id" form:"id" gorm:"primary_key"`
	OrderSn  string          `json:"order_sn" form:"order_sn"`
	UserId   int             `json:"user_id" form:"user_id"`
	Phone    string          `json:"phone" form:"phone"`
	Amount   decimal.Decimal `json:"amount" form:"amount"`
	Integral int             `json:"integral" form:"integral"`
	Details  []Detail        `json:"detail" form:"detail" gorm:"-"`
}

type Detail struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
	Sn   string `json:"sn"`
}

// TableName 定义表明
func (u *Order) TableName() string {
	return "order"
}
func (d *Detail) TableName() string {
	return "detail"
}

func Add(c *gin.Context) {
	var order Order
	var book []bo.Book
	var detail []Detail
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
	}
	timestamp := time.Now().Unix()
	random := rand.Intn(10000)
	order.OrderSn = fmt.Sprintf("%d%d", timestamp, random) //string(timestamp) + string(random)
	slice := make([]Detail, 0)
	global.GlobalDB.Table("book").Find(&book)
	for _, v := range order.Details {
		slice = append(slice, Detail{
			Name: v.Name,
			Num:  v.Num,
			Sn:   order.OrderSn,
		})
	}
	accumulator := decimal.NewFromInt(0)
	for k, v := range book {
		for i := 0; i < len(slice); i++ {
			if book[k].BookName == slice[i].Name {
				accumulator = accumulator.Add(v.Money.Mul(decimal.NewFromInt(int64(slice[i].Num))))
			}
		}
	}
	order.Amount = accumulator
	detail = order.Details
	for i := 0; i < len(detail); i++ {
		detail[i].Sn = order.OrderSn
	}
	tx := global.GlobalDB.Begin() //开启事务
	tx.Create(&order)
	tx.Table("order_details").CreateInBatches(&detail, len(detail))
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // 失败时回滚事务
		c.JSON(http.StatusOK, common.DbFail(err.Error()))
		return
	}
}
