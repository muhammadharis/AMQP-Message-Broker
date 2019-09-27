package main

import (
	"fmt"
	"context"
	"io"

	consumer "github.com/muhammadharis/grpc/protos/consumer"
)

// Subscribe allows a consumer to consume a message from the queue
func Subscribe(key string) error {
	conn, err := CreateGrpcClientConnection("localhost:8080")
	if err != nil {
		return err
	}
	client := consumer.NewConsumerAPIClient(conn)

	consumerRequest := &consumer.ConsumerRequest{
		Key: key,
	}

	stream, err := client.Subscribe(context.Background(), consumerRequest)
	if err != nil {
		return err
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(msg.GetMessage())
	}
}
