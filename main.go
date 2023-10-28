package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"net/http"
	"time"
	"wwgo/router"
)

var start = time.Now()

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
	router.Router(r)
	//创建调度器
	s := gocron.NewScheduler()
	s.Every(1).Second().Do(task) //设置频率和执行的方法
	//s.Start()
	defer s.Clear()
	//3.监听端口
	r.Run(":8075")
}
func task() {
	time := time.Now()
	fmt.Println("程序已经运行的时间：", time.Sub(start)) //格式化时间戳
}
