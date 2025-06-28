package db

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func TestQuickStart() {
	// 创建Redis连接客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.10.115:6379",
		Password: "Dingtalk1234561017",
		DB:       0, // 使用默认DB
	})

	// 设置键值对，0就是永不过期
	redisClient.Set("hello", "world", 0)

	// 读取值
	result, err := redisClient.Get("hello").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("ket not exist")
	} else if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)
}
