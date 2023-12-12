package config

import (
	"github.com/go-redis/redis"
	"log"
)

func GetRedis() *redis.Client {
	config, err := GetConfig()
	if err != nil {
		log.Fatalf("fail to get config: %v", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Save().Err()
	if err != nil {
		log.Fatal("持久化失败")
	}
	return rdb
}
