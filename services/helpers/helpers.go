package helpers

import (
	redis "github.com/go-redis/redis"
)

func CreateRedisClient(addr string, password string, db int) (*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}

func createGrpcClientConnection(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}