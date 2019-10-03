package main

import (
	"fmt"
	"context"
	"io"

	broker "github.com/muhammadharis/AMQP-Message-Broker/protos/broker"
)

// InitializeConsumerClient initializes the Consumer API client
func InitializeConsumerClient(host string) (broker.ConsumerAPIClient){
	conn, err := openConnection(host)
	if err != nil {
		panic(err)
	}
	// Initialize the client
	client := broker.NewConsumerAPIClient(conn)
	return client
}

// CreateReadStream opens a stream that the consumer can read messages from
func CreateReadStream(client broker.ConsumerAPIClient) (broker.ConsumerAPI_ReadMessageClient, error) {
	consumerRequest := &broker.ConsumerReadRequest{
		QueueName: "my_queue",
	}
	return client.ReadMessage(context.Background(), consumerRequest)
}

// Read allows a consumer to consume a message from the queue
func Read() error {

	// Initialize client
	client := InitializeConsumerClient("localhost:8080")

	// Create a reading stream
	stream, err := CreateReadStream(client)
	if err != nil {
		return err
	}

	// Receive messages on the stream
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(msg.GetMessage()) //Print the messages
	}
}

