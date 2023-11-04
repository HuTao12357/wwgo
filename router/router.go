package router

import (
	"github.com/gin-gonic/gin"
	"wwgo/api/bo"
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
	}
	//书
	book := r.Group("book")
	{
		book.GET("getById", bo.GetById)
	}
}
