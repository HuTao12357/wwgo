package us

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"Username"`
	Password string `json:"password" form:"password"`
}

// Login 登录
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
	var Dword string
	var id string
	db := getNewDB()
	//sql := fmt.Sprintf("select * from user where username='%s' ", user.Username) //原生sql
	result := db.Table("user").Where("username = ?", name).First(&user)
	if result.RowsAffected == 0 {
		fmt.Println("登陆查询数据库没有数据")
	} else {
		Dword = user.Password
	}
	a := utils.ComparePasswords(Dword, []byte(password))
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
		token, err := utils.GenToken(name, id)
		if err != nil {
			fmt.Println("生成token失败")
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"msg":   "登陆成功",
			"token": token,
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

// InsertOrUpdate 新增更新
func InsertOrUpdate(c *gin.Context) {
	var user User
	db := getNewDB()
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Print(err)
	}
	if user.Id != 0 { //新增
		//密码加密
		word := []byte(user.Password)
		password := utils.HashAndSalt(word)

		newData := User{
			Username: user.Username,
			Password: password,
		}
		db.Table("user").Create(newData)
	} else {
		word := []byte(user.Password)
		password := utils.HashAndSalt(word)
		user.Password = password
		db.Table("user").Updates(user)
	}

}

// GetById id查询
func GetById(c *gin.Context) {
	id := c.Query("id") //接收前端传的id
	var user User
	db := getNewDB()
	//通过Debug模式查看执行的sql
	res := db.Debug().Table("user").Where("id=?", id).Find(&user)
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

// UsersQuery 列表查询
func UsersQuery(c *gin.Context) {
	var users []User
	PageSize, _ := strconv.ParseInt(c.Query("PageSize"), 10, 64)
	PageNo, _ := strconv.ParseInt(c.Query("PageNo"), 10, 64)
	offset := (PageNo - 1) * PageSize
	db := getNewDB()
	db.Table("user").Offset(int(offset)).Limit(int(PageSize)).Order("id DESC").Find(&users)
	//对分页后的数据进行操作
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": users,
	})
	return
}

// BatchUserInsert 批量添加,让前端传json
func BatchUserInsert(c *gin.Context) {
	var user []User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "无效参数",
		})
		return
	}
	db := getNewDB()
	for i := range user {
		user[i].Password = utils.HashAndSalt([]byte(user[i].Password))
	}
	res := db.Table("user").CreateInBatches(&user, len(user))
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "批量添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "批量添加成功",
	})
	return
}

// GetEnrollNum @Summary	查询每月的注册人数
// @Produce json
// @Param year query string false "年份"
// @Success 200	{object} string "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /User/GetEnrollNum [get]
func GetEnrollNum(c *gin.Context) {
	data := c.Query("year")
	fmt.Println(data)
	db := getNewDB()

	var monthNum []map[string]interface{}
	var sql = fmt.Sprintf("SELECT months.month, IFNULL(data.count, 0) AS count\nFROM (\n  SELECT 1 AS month UNION SELECT 2 UNION SELECT 3 UNION SELECT 4\n  UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8\n  UNION SELECT 9 UNION SELECT 10 UNION SELECT 11 UNION SELECT 12\n) AS months\nLEFT JOIN (\n  SELECT MONTH(update_at) AS month, COUNT(*) AS count\n  FROM user as a \n  GROUP BY month\n) AS data ON months.month = data.month\nORDER BY months.month")
	db.Raw(sql).Scan(&monthNum)
	for k := range monthNum { //返回键和值,  迭代返回的值只是映射一个副本，而不是原始映射
		if monthNum[k]["month"].(int64) == 1 { //interface {} is int64, not int
			monthNum[k]["first"] = "第一月"
		}
	}
	fmt.Println("====", monthNum)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "响应成功",
		"data": monthNum,
	})
	return
}
