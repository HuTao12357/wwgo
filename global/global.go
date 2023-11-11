package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"sync"
	"time"
)

var (
	GlobalDB *gorm.DB
	once     sync.Once //保证只初始化一次
)

var dbName = os.Getenv("DB_NAME")
var dbPassword = os.Getenv("DB_PASSWORD")
var err error

const port = "3306"

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/wwgo?charset=utf8mb4&parseTime=True&loc=Local", dbName, dbPassword, port)
	GlobalDB, err = gorm.Open(mysql.New(mysql.Config{ //不要使用：=,不然数据库的连接只会在局部变量中生效，要赋值给全局变量
		DSN:               dsn,
		DefaultStringSize: 171, //String类型的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction: false, //不跳过默认默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gorm_", //前缀
			SingularTable: true,    //使用单数表名
		},
	})
	if err != nil {
		// 处理连接错误
		fmt.Println("连接错误==================")
	}
	/*
		gorm库中没有Close()方法,因为DB实例在创建时被存储在一个连接池中，需要的时候自动获取连接，因此在使用完毕之后不用主动关闭数据库连接
	*/
	sqlDB, _ := GlobalDB.DB()           //GORM使用database/sql 来维护连接池
	sqlDB.SetMaxIdleConns(10)           //连接池最大的空闲连接数
	sqlDB.SetMaxOpenConns(40)           //连接池最大连接数量
	sqlDB.SetConnMaxIdleTime(time.Hour) //连接池中连接的最大可复用时间
}
