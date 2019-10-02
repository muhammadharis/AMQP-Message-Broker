package helpers

import (
	"os"
	"github.com/go-redis/redis"
)

const (
	ExchangeNamespace = "_EXCHANGE_"
	QueueNamespace = "_QUEUE_"
	StreamNamespace = "_STREAM_"
)

// CreateRedisClient creates and returns a Redis client with given information
func CreateRedisClient(password string, db int) (*redis.Client){
	addr := "localhost:6379" //Default host and port
	if os.Getenv("REDIS_HOST") != "" && os.Getenv("REDIS_PORT") != "" {
		addr = os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}