package main

import (
	//"fmt"
	"net"
	//"time"

	broker "github.com/muhammadharis/AMQP-Message-Broker/protos/broker"
	brokerImpl "github.com/muhammadharis/AMQP-Message-Broker/services/broker"
	consumerImpl "github.com/muhammadharis/AMQP-Message-Broker/services/consumer"
	//redis "github.com/go-redis/redis"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	broker.RegisterProduceAPIServer(server, &brokerImpl.BrokerImpl{})
	broker.RegisterConsumerAPIServer(server, &consumerImpl.ConsumerImpl{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server.Serve(listener)
	/*client := redis.NewClient(&redis.Options{
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
	fmt.Println(val+1)*/
}