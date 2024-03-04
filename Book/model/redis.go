package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func SetRedis(key string, value any, mytime time.Duration) error {
	//设置键值对：使用Set方法设置键为"key"，值为"value"的数据，并设置过期时间为0
	err := ConnectRedis.Set(context.Background(), key, value, mytime).Err()
	return err
}
func GetRedis(key string) string {
	//获取键的值：使用Get方法获取键为"key"的值，并根据返回结果进行判断。
	value, err := ConnectRedis.Get(context.Background(), "key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return ""
	} else {
		fmt.Println("key:", value)
	}
	return value
}
func DelRedis(key string) error {
	//删除键值对
	err := ConnectRedis.Del(context.Background(), key).Err()
	return err
}
