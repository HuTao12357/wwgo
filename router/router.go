package router

import (
	"github.com/gin-gonic/gin"
	"wwgo/api/us"
)

func Router(r *gin.Engine) {
	//用户
	user := r.Group("User")
	{
		user.POST("login", us.Login)
	}
}
