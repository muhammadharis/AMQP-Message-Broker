package helpers

import (
	"github.com/go-redis/redis"
)


func CreateRedisClient(addr string, password string, db int) (*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}