package v1

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "go-gin-dcard/middleware"
    "go-gin-dcard/models"
)


/**
* 配對卡友
* @method GET
* @router /api/v1/limit/pair
*/
func GetPairApi(c *gin.Context) {
    ip := c.ClientIP()
    p := models.Pair{Ip: ip}

    // do pair
    reqTimes, err := p.DoPair()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    remainingReqTimes := middleware.RateLimitTimes - reqTimes

    c.Writer.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(remainingReqTimes, 10))
    c.Writer.Header().Set("X-RateLimit-Reset", c.MustGet("ttl").(string))
    c.JSON(http.StatusOK, gin.H{
        "result": reqTimes,
        "msg": "Successful pair cards!",
    })
}