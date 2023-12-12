package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// Config 创建结构体解析配置信息
type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	}
}

func GetConfig() (*Config, error) {
	filepath := "config.yml"
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("fail to read config file: %v", err)
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(file, config) //将yml数据解析目标结构体
	if err != nil {
		log.Fatalf("file to Unmarshal file :%v", err)
		return nil, err
	}
	return config, nil
}
func MysqlGet() (*gorm.DB, error) {
	config, err := GetConfig()
	if err != nil {
		log.Fatalf("fail to get config: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to connect database :%v", err)
	}
	return db, err
}
