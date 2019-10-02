package main

import (
	"context"

	broker "github.com/muhammadharis/grpc/protos/broker"
	"google.golang.org/grpc"
)

func openConnection(host string) (*grpc.ClientConn, error) {
	conn, err := CreateGrpcClientConnection("localhost:8080")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// InitializeProducerClient initializes the Produce API client
func InitializeProducerClient(host string) (broker.ProduceAPIClient){
	conn, err := openConnection(host)
	if err != nil {
		panic(err)
	}
	// Initialize the client
	client := broker.NewProduceAPIClient(conn)
	return client
}

// CreateExchange creates a new exchange with a specified name
func CreateExchange(client broker.ProduceAPIClient, exchangeName, exchangeType string) {
	createExchangeRequest := &broker.CreateExchangeRequest {
		ExchangeName : exchangeName,
		Type: exchangeType,
	}
	client.CreateExchange(context.Background(), createExchangeRequest)
}

// CreateQueue creates a new Queue with a specified name
func CreateQueue(client broker.ProduceAPIClient, queueName string) {
	createQueueRequest := &broker.CreateQueueRequest {
		QueueName: queueName,
	}
	client.CreateQueue(context.Background(), createQueueRequest)
}

// CreateBinding binds an existing queue to an existing exchange
func CreateBinding(client broker.ProduceAPIClient, exchangeName, queueName string) {
	bindingRequest := &broker.BindQueueRequest {
		ExchangeName: exchangeName,
		QueueName: queueName,
	}
	client.BindQueue(context.Background(), bindingRequest)
}

// ProduceMessage sends a message to the correct queue and exchange
func ProduceMessage(client broker.ProduceAPIClient, exchangeName, queueName string, messageSet [][]byte){
	produceRequest := &broker.ProduceRequest {
		ExchangeName: exchangeName,
		QueueName: queueName,
		MessageSet: [][]byte{[]byte("Hello, world!"), []byte("This is a message")},
	}
	client.ProduceMessage(context.Background(), produceRequest)
}

// Message allows a producer to produce a fanout message to the broker
func Message() error {
	// Initialize the client
	client := InitializeProducerClient("localhost:8080")

	// Create a new exchange
	CreateExchange(client, "my_exchange", "fanout")

	// Create a new queue
	CreateQueue(client, "my_queue")

	// Create a binding between queue and exchange
	CreateBinding(client, "my_exchange", "my_queue")

	// Produce a direct message to my_queue
	ProduceMessage(client, "my_exchange", "my_queue", [][]byte{[]byte("Hello, world!"), []byte("This is a message")})
	
	return nil
}