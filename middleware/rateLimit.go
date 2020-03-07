package middleware

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    
    "go-gin-dcard/database/redis"
    "go-gin-dcard/models"
)

var RateLimitTimes int64 = 1000

/* Limit the number of requests from the same IP per hour less than 1000 */
func RateLimitMiddle() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP() // get ip address
        p := models.Pair{Ip: ip}

        /* check does the ip's key exist */
        isExist, err := p.CheckIpKeyExist()
        if err != nil {
            c.JSON(http.StatusServiceUnavailable, gin.H{
                "result": "",
                "msg": err.Error(),
            })
            return
        }

        // if there is no exist this key in redis, insert one
        if !isExist {
            err = p.SetNewOne()
            if err != nil {
                c.JSON(http.StatusServiceUnavailable, gin.H{
                    "result": "",
                    "msg": err.Error(),
                })
                return
            }
        }

        // get the key remaining time in redis
        ttl, err := redis.RedisClient.TTL(ip).Result()
        if err != nil {
            return
        }
        
        /* check does this ip over the limit of the number of requests */
        // get req times of the ip
        reqTimes, err := redis.RedisClient.Get(ip).Result()
        if err != nil {
            return
        }
        
        // if over the limit of the number of requests, return 429 TooManyRequests
        reqTimesInt64, _ := strconv.ParseInt(reqTimes, 10, 64)
        if reqTimesInt64 == RateLimitTimes {
            c.Writer.Header().Set("X-RateLimit-Remaining", "0")
            c.Writer.Header().Set("X-RateLimit-Reset", ttl.String())

            c.JSON(http.StatusTooManyRequests, gin.H{
                "result": "",
                "msg": "Over the limit of the number of requests",
            })

            c.Abort()
            return
        }
        
        c.Set("ttl", ttl.String()) // set ttl in context
        c.Next()
        return
    }
}