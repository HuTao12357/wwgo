package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var err error

var dbName = os.Getenv("DB_NAME")
var dbPassword = os.Getenv("DB_PASSWORD")

const port = "3306"

func GetMysql() (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/wwgo?charset=utf8mb4&parseTime=True&loc=Local", dbName, dbPassword, port)
	db, err = gorm.Open(mysql.New(mysql.Config{
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
	sqlDB, _ := db.DB()                 //GORM使用database/sql 来维护连接池
	sqlDB.SetMaxIdleConns(10)           //连接池最大的空闲连接数
	sqlDB.SetMaxOpenConns(40)           //连接池最大连接数量
	sqlDB.SetConnMaxIdleTime(time.Hour) //连接池中连接的最大可复用时间
	return
}
func init() {
	GetMysql()
}
