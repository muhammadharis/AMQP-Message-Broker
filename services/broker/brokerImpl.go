package broker

import (
	"fmt"
	"context"
	"errors"
	"strings"
	broker "github.com/muhammadharis/grpc/protos/broker"
	redis "github.com/go-redis/redis"
	
	helpers "github.com/muhammadharis/grpc/services/helpers"
)

const (
	partitionLimit = 10
	exchangeNamespace = "_EXCHANGE_"
	queueNamespace = "_QUEUE_"
	streamNamespace = "_STREAM_"
)

// BrokerImpl implements the server code for the Broker proto
type BrokerImpl struct {}

// CreateExchange creates a new exchange with a specified name
func (*BrokerImpl) CreateExchange(ctx context.Context, request *broker.CreateExchangeRequest) (*broker.CreateExchangeResponse, error) {
	exchangeName := request.ExchangeName
	exchangeType := strings.ToUpper(request.Type)
	if exchangeType == "" { // The default exchange is a fanout exchange
		exchangeType = "FANOUT"
	}

	client := helpers.CreateRedisClient("localhost:6379", "", 0)

	exists, err := client.Exists(exchangeNamespace+exchangeName).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	if (exists > 0) {
		panic(errors.New("An exchange with the name "+exchangeName+" already exists."))
		return  nil, errors.New("An exchange with the name "+exchangeName+" already exists.")
	}

	// Create a key that maps the exchange name to its binded queues
	// Add the type of the exchange into the map with the hash "TYPE"
	boolCmd := client.HSet(exchangeNamespace+exchangeName, "TYPE", exchangeType)
	if boolCmd.Err() != nil {
		panic(boolCmd.Err())
		return nil, boolCmd.Err()
	}

	return &broker.CreateExchangeResponse{}, nil
}

// CreateQueue creates a new Queue with a specified name
func (*BrokerImpl) CreateQueue(ctx context.Context, request *broker.CreateQueueRequest) (*broker.CreateQueueResponse, error) {
	queueName := request.QueueName

	client := helpers.CreateRedisClient("localhost:6379", "", 0)

	exists, err := client.Exists(queueNamespace+queueName).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	if (exists > 0) {
		panic(errors.New("A queue with the name "+queueName+" already exists."))
		return  nil, errors.New("A queue with the name "+queueName+" already exists.")
	}

	// Create a key that maps the queue name to its binded exchanges
	// The default exchange named DEFAULT is binded to all queues
	boolCmd := client.HSet(queueNamespace+queueName, "DEFAULT", true)
	if boolCmd.Err() != nil {
		panic(boolCmd.Err())
		return nil, boolCmd.Err()
	}

	return &broker.CreateQueueResponse{}, nil
}

// BindQueue binds an existing queue to an existing exchange
func (*BrokerImpl) BindQueue(ctx context.Context, request *broker.BindQueueRequest) (*broker.BindQueueResponse, error) {
	exchangeName := request.ExchangeName
	queueName := request.QueueName
	
	client := helpers.CreateRedisClient("localhost:6379", "", 0)

	// Add the specified queue to the set of queues held by the specified exchange, and vice versa
	cmd := client.HSet(exchangeNamespace+exchangeName, queueName, true)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	cmd = client.HSet(queueNamespace+queueName, exchangeName, true)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	
	return &broker.BindQueueResponse{}, nil
}

// Produce consumes a producer's message
func (*BrokerImpl) ProduceMessage(ctx context.Context, request *broker.ProduceRequest) (*broker.ProduceResponse, error) {
	exchangeName := request.ExchangeName

	client := helpers.CreateRedisClient("localhost:6379", "", 0) //No password, default DB

	exchangeType, err := client.HGet(exchangeNamespace+exchangeName, "TYPE").Result() //Gets the type of the exchange
	if err != nil {
		panic(err)
		return nil, err
	}
	exchangeType = strings.ToUpper(exchangeType)
	fmt.Println(exchangeType)
	
	switch exchangeType {
		case "FANOUT":
			err = produceFanoutHelper(client, request)
		case "DIRECT":
			err = produceDirectHelper(client, request)
	}
/*
	xAddArgument := &redis.XAddArgs{
		Stream: streamName,
		MaxLen: partitionLimit,
		MaxLenApprox: partitionLimit,
		ID: "*",
		Values: map[string]interface{} {
			"0": request.MessageSet,
		},
	}
	cmd := client.XAdd(xAddArgument)
	fmt.Println(cmd.String())
*/
	resp := &broker.ProduceResponse{}
	return resp, nil
}

func produceFanoutHelper(client *redis.Client, request *broker.ProduceRequest) error {
	return nil
}

func produceDirectHelper(client *redis.Client, request *broker.ProduceRequest) error {
	queueName := request.QueueName
	xAddArgument := &redis.XAddArgs{
		Stream: streamNamespace+queueName,
		MaxLen: partitionLimit,
		MaxLenApprox: partitionLimit,
		ID: "*",
		Values: map[string]interface{} {
			"0": request.MessageSet,
		},
	}
	cmd := client.XAdd(xAddArgument)
	fmt.Println(cmd.String())
	return nil
}
