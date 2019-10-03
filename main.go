package main

import (
	"net"

	broker "github.com/muhammadharis/AMQP-Message-Broker/protos/broker"
	brokerImpl "github.com/muhammadharis/AMQP-Message-Broker/services/broker"
	consumerImpl "github.com/muhammadharis/AMQP-Message-Broker/services/consumer"

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
}