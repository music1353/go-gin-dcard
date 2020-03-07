package router

import (
    "github.com/gin-gonic/gin"

    "go-gin-dcard/middleware"
    "go-gin-dcard/router/api/v1"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    apiv1 := r.Group("/api/v1")
    {
        limit := apiv1.Group("/limit", middleware.RateLimitMiddle())
        {
            limit.GET("/pair", v1.GetPairApi)
        }
    }
    
    return r
}