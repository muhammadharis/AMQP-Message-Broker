package main

import (
	"fmt"
	"context"
	"io"

	broker "github.com/muhammadharis/AMQP-Message-Broker/protos/broker"
)

// Read allows a consumer to consume a message from the queue
func Read() error {
	conn, err := CreateGrpcClientConnection("localhost:8080")
	if err != nil {
		return err
	}
	client := broker.NewConsumerAPIClient(conn)

	consumerRequest := &broker.ConsumerReadRequest{
		QueueName: "my_queue",
	}
	
	stream, err := client.ReadMessage(context.Background(), consumerRequest)
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

