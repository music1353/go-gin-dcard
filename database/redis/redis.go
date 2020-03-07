package redis

import (
    "log"

    "github.com/go-redis/redis/v7"
        
    "go-gin-dcard/pkg/setting"
)

var RedisClient *redis.Client

func Setup() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     setting.RedisSetting.Host,
        Password: setting.RedisSetting.Password,
        DB:       setting.RedisSetting.DB,
    })

    _, err := RedisClient.Ping().Result()
    if err != nil {
        log.Fatal("Connect to Redis error:", err.Error())
    }
    log.Println("Successful connected to Redis")
}

func Exist(key string) (bool, error) {
    _, err := RedisClient.Get(key).Result()
	if err == redis.Nil {
        return false, nil
    } else if err != nil {
        return false, err
    }

    return true, nil
}