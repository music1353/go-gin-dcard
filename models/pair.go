package models

import (
    "time"
    
    "go-gin-dcard/database/redis"
)

type Pair struct {
    Ip string // the ip of dcard user who requests to pair
}

/**
* redis內是否已有此ip的cache
* @param {string} ip - dcard user's ip
* @return {bool} does the key exist
* @return {error}
*/
func (p *Pair) CheckIpKeyExist() (isExist bool, err error) {
    isExist, err = redis.Exist(p.Ip)
    if err != nil {
        return
    }
    return
}

func (p *Pair) SetNewOne() (err error) {
    err = redis.RedisClient.Set(p.Ip, 0, 1*time.Hour).Err()
    if err != nil {
        return
    }
    return nil
}

/**
* 配對卡有
* @param {string} ip - dcard user's ip
* @return {int32}
* @return {error}
*/
func (p *Pair) DoPair() (reqTimes int64, err error) {
    /* business logic of pair card
        ...
        ...
    */
    
    reqTimes, err = redis.RedisClient.Incr(p.Ip).Result()
    if err != nil {
        return
    }
    return
}

