package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wwgo/config"
)

type Menu struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Pid      int    `json:"pid"`
	Children []Menu `json:"children" gorm:"-"`
}

func GetMenuTree(c *gin.Context) {
	var menuPID []Menu
	var menuZI []Menu
	db, _ := config.MysqlGet()
	db.Table("menu").Where("pid = 0").Find(&menuPID) // Only select the necessary class A fields
	db.Table("menu").Where("pid != 0").Find(&menuZI) // Only select the necessary class Not A fields
	var all []interface{}
	for _, v := range menuPID {
		menus := iterateMenu(menuZI, v.Id)
		v.Children = menus
		fmt.Println(v)
		all = append(all, v)
	}
	c.JSON(http.StatusOK, all)
}

func iterateMenu(zi []Menu, id int) []Menu {
	var me []Menu
	for _, v := range zi {
		if v.Pid == id {
			s := iterateMenu(zi, v.Id) // Use Children for recursion
			v.Children = s
			me = append(me, v)
		}
	}
	return me
}
