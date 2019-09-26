package main

import (
	"fmt"
	"context"

	broker "github.com/muhammadharis/grpc/protos/broker"

	"google.golang.org/grpc"
)

func createGrpcClientConnection(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Produce allows a producer to produce a message to the broker
func Produce(routingKey string) {
	conn, err := createGrpcClientConnection("localhost:8080")
	if err != nil {
		panic(err)
	}
	client := broker.NewProduceAPIClient(conn)

	produceRequest := &broker.ProduceRequest{
		RoutingKey: routingKey,
		MessageSet: []byte("example message set"),
	}

	produceResponse, err := client.Produce(context.Background(), produceRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println(produceResponse)
}