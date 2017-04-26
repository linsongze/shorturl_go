package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/linsongze/shorturl_go/config"
	"time"
)
type RedisStoreService struct {
	RedisClient     *redis.Pool
}

func NewRedisStoreService()*RedisStoreService {
	rs :=  new(RedisStoreService)
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := config.Redis_host+":"+config.Redis_port
	// 建立连接池
	rs.RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     100,
		MaxActive:   10,
		IdleTimeout: 360 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	// 从池里获取连接
	rc := rs.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	rc.Do("set","sid",0)
	return rs
}
func (rs *RedisStoreService) Save(shortCode, url string) {
	// 从池里获取连接
	rc := rs.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_,err := rc.Do("HSET","ss",shortCode,url)
	if err !=nil{
		fmt.Println(err)
	}
}
func (rs *RedisStoreService) IncAndGet() int64 {
	// 从池里获取连接
	rc := rs.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	val,err := rc.Do("INCR","sid")
	if err !=nil{
		fmt.Println(err)
		return 0
	}
	return val.(int64)
}
func (rs *RedisStoreService) Get(shortCode string) string {
	// 从池里获取连接
	rc := rs.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	val,err := redis.String(rc.Do("HGET","ss",shortCode))
	if err !=nil{
		fmt.Println(err)
	}
	return val
}