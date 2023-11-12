package global

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"sync"
	"time"
)

var (
	GLOBAL_PGDB *gorm.DB
	first       sync.Once
)
var PGName = "postgres"
var PGWord = os.Getenv("DB_PASSWORD")

const PGport = "5432"

func init() {
	dsn := fmt.Sprintf("host=127.0.0.1 port=%s user=%s password=%s dbname=wwgo sslmode=disable TimeZone=Asia/Shanghai", PGport, PGName, PGWord)

	GLOBAL_PGDB, err = gorm.Open(postgres.New(postgres.Config{ //不要使用：=,不然数据库的连接只会在局部变量中生效，要赋值给全局变量
		DSN:                  dsn,
		PreferSimpleProtocol: true,
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
	sqlDB, _ := GLOBAL_PGDB.DB()        //GORM使用database/sql 来维护连接池
	sqlDB.SetMaxIdleConns(10)           //连接池最大的空闲连接数
	sqlDB.SetMaxOpenConns(40)           //连接池最大连接数量
	sqlDB.SetConnMaxIdleTime(time.Hour) //连接池中连接的最大可复用时间
}
