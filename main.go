package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "hello world",
		})
	})
	//3.监听端口
	r.Run(":8075")
}
