package main

import (
	"context"

	broker "github.com/muhammadharis/grpc/protos/broker"
)

// Message allows a producer to produce a fanout message to the broker
func Message() error {
	conn, err := CreateGrpcClientConnection("localhost:8080")
	if err != nil {
		return err
	}
	
	// Initialize the client
	client := broker.NewProduceAPIClient(conn)

	// Create a new exchange
	createExchangeRequest := &broker.CreateExchangeRequest {
		ExchangeName : "my_exchange",
  		Type: "fanout",
	}
	client.CreateExchange(context.Background(), createExchangeRequest)

	// Create a new queue
	createQueueRequest := &broker.CreateQueueRequest {
		QueueName: "my_queue",
	}
	client.CreateQueue(context.Background(), createQueueRequest)

	// Create a binding between queue and exchange
	bindingRequest := &broker.BindQueueRequest {
		ExchangeName: "my_exchange",
		QueueName: "my_queue",
	}
	client.BindQueue(context.Background(), bindingRequest)

	// Produce a direct message to my_queue
	produceRequest := &broker.ProduceRequest {
		ExchangeName: "my_exchange",
		QueueName: "my_queue",
		MessageSet: [][]byte{[]byte("Hello, world!"), []byte("This is a message")},
	}
	client.ProduceMessage(context.Background(), produceRequest)
	return nil
}