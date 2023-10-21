package us

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"wwgo/connection"
	"wwgo/utils"
)

func getNewDB() (D *gorm.DB) {
	return connection.GetMysql()
}

type User struct {
	Id       string `json:"id" form:"id"`
	Username string `json:"username" form:"Username"`
	Password string `json:"password" form:"password"`
}

// 登录
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
	var DBword string
	db := getNewDB()
	//sql := fmt.Sprintf("select * from user where username='%s' ", user.Username) //原生sql
	result := db.Table("user").Where("username = ?", name).First(&user)
	if result.RowsAffected == 0 {
		fmt.Println("登陆查询数据库没有数据")
	} else {
		DBword = user.Password
	}
	a := utils.ComparePasswords(DBword, []byte(password))
	if a == false {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "账号或密码不正确",
		})
		return
	}

	var count int64
	count = db.Table("user").Where("username = ?", name).Where("password = ?", password).Count(&count).RowsAffected
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "登陆成功",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "登陆失败",
		})
		return
	}
}

// 新增
func Insert(c *gin.Context) {
	var user User
	db := getNewDB()
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Print(err)
	}
	//密码加密
	word := []byte(user.Password)
	password := utils.HashAndSalt(word)
	ids := uuid.NewV4().String()
	newData := User{
		Username: user.Username,
		Password: password,
		Id:       ids,
	}
	db.Table("user").Create(newData)
}

// id查询
func GetById(c *gin.Context) {
	id := c.Query("id") //接收前端传的id
	var user User
	db := getNewDB()
	res := db.Table("user").Where("id=?", id).First(&user)
	if res.Error != nil {
		fmt.Println("=======查询数据库出错")
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "该用户不存在,请先注册",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": user,
		})
		return
	}
}

// 列表查询
func UsersQuery(c *gin.Context) {
	var users []User
	PageSize, _ := strconv.ParseInt(c.Query("PageSize"), 10, 64)
	PageNo, _ := strconv.ParseInt(c.Query("PageNo"), 10, 64)
	offset := (PageNo - 1) * PageSize
	db := getNewDB()
	db.Table("user").Offset(int(offset)).Limit(int(PageSize)).Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": users,
	})
	return
}
