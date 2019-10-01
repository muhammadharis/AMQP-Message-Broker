package broker

import (
	"fmt"
	"context"
	"errors"
	//"strconv"
	broker "github.com/muhammadharis/grpc/protos/broker"
	redis "github.com/go-redis/redis"
	
	helpers "github.com/muhammadharis/grpc/services/helpers"
)

const (
	partitionLimit = 10
	exchangeNamespace = "_EXCHANGE_"
	queueNamespace = "_EXCHANGE_"
)

type BrokerImpl struct {}

// CreateExchange creates a new exchange with a specified name
func (*BrokerImpl) CreateExchange(ctx context.Context, request *broker.CreateExchangeRequest) (*broker.CreateExchangeResponse, error) {
	exchangeName := request.ExchangeName
	exchangeType := request.Type
	if exchangeType == "" { // The default exchange is a fanout exchange
		exchangeType = "fanout"
	}

	client := helpers.CreateRedisClient("localhost:6379", "", 0)

	exists, err := client.Exists(exchangeNamespace+exchangeName).Result()
	if err != nil {
		return nil, err
	}
	if (exists > 0) {
		return  nil, errors.New("An exchange with the name "+exchangeName+" already exists.")
	}

	cmd := client.Set(exchangeNamespace+exchangeName, map[int]bool{}, 0)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	return &broker.CreateExchangeResponse{}, nil
}

// CreateQueue creates a new Queue with a specified name
func (*BrokerImpl) CreateQueue(ctx context.Context, request *broker.CreateQueueRequest) (*broker.CreateQueueResponse, error) {
	queueName := request.QueueName

	client := helpers.CreateRedisClient("localhost:6379", "", 0)

	exists, err := client.Exists(queueNamespace+queueName).Result()
	if err != nil {
		return nil, err
	}
	if (exists > 0) {
		return  nil, errors.New("A queue with the name "+queueName+" already exists.")
	}

	cmd := client.Set(queueNamespace+queueName, map[int]bool{}, 0)
	if cmd.Err() != nil {
		return nil, cmd.Err()
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
func (*BrokerImpl) Produce(ctx context.Context, request *broker.ProduceRequest) (*broker.ProduceResponse, error) {
	streamName := request.StreamName
	//messageSet := request.MessageSet
	client := helpers.CreateRedisClient("localhost:6379", "", 0) //No password, default DB
	
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

	resp := &broker.ProduceResponse{}
	return resp, nil
}

