package main

import (
	"fmt"
	//"net"
	//"time"

	broker "github.com/muhammadharis/grpc/protos/broker"
	brokerImpl "github.com/muhammadharis/grpc/services/broker"
	redis "github.com/go-redis/redis"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	broker.RegisterProduceAPIServer(server, &brokerImpl.BrokerImpl{})

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client.LPush("key", "a", "b", "c", "d")
	
	fmt.Println(client.LRange("key", 0, -1))

	client.Set("key1", "52", 0)
	val, err := client.Get("key1").Int64()
	if err != nil {
		panic(err)
	}
	fmt.Println(val+1)
}