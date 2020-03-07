package main

import (
    "github.com/gin-gonic/gin"
    "go-gin-dcard/router"

    "go-gin-dcard/pkg/setting"
    "go-gin-dcard/database/redis"
)

func init() {
    setting.Setup()
    redis.Setup()
}

func main() {
    gin.SetMode(setting.ServerSetting.RunMode)

    router := router.InitRouter()
    router.Run()
}