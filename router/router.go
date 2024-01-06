package router

import (
	"github.com/gin-gonic/gin"
	"wwgo/api/bo"
	"wwgo/api/order"
	"wwgo/api/us"
)

func Router(r *gin.Engine) {
	//用户
	user := r.Group("User")
	{
		user.POST("login", us.Login)
		user.POST("InsertOrUpdate", us.InsertOrUpdate)
		user.GET("GetById", us.GetById)
		user.GET("UserQuery", us.UsersQuery)
		user.POST("BatchUserInsert", us.BatchUserInsert)
		user.GET("GetEnrollNum", us.GetEnrollNum)
		user.POST("ExecInsert", us.ExecInsert)
		user.POST("inGet", us.InGet)
	}
	//书
	book := r.Group("book")
	{
		book.GET("getById", bo.GetById)
		book.GET("getRate", bo.GetNumRate)
		book.POST("getExcel", bo.GetExcel)
		book.GET("testResult", bo.TestResult)
	}
	or := r.Group("order")
	{
		or.POST("add", order.Add)
		or.POST("page", order.Page)
		or.GET("GetDetail", order.GetDetail)
	}

}
