package broker

import (
	"context"
	broker "github.com/muhammadharis/grpc/protos/broker"
	redis "github.com/go-redis/redis"
)

const partitionLimit = 10

type BrokerImpl struct {}

// Produce consumes a producer's message
func (*BrokerImpl) Produce(ctx context.Context, request *broker.ProduceRequest) (*broker.ProduceResponse, error) {
	routingKey := request.RoutingKey
	messageSet := request.MessageSet
	client := createRedisClient("localhost:6379", "", 0) //No password, default DB
	dirtyPartition := false // Is set when the partition number is changed

	// The routing key will hold the partition number, actual data is stored in routingkey+partitionNumber
	partitionNumber, err := client.Get(routingKey).Int64()
	if err != nil {
		panic(err)
	}

	sizeOfCurrentPartition := client.LLen(routingKey+string(partitionNumber)).Val()
	// Increment the partition number if the partition is at its max capacity
	if sizeOfCurrentPartition >= partitionLimit {
		dirtyPartition = true
		partitionNumber++
		sizeOfCurrentPartition = 0
	}

	for msg := range messageSet {
		client.RPush(routingKey+string(partitionNumber), msg) //We store messages in routingKey+partitionNumber
		sizeOfCurrentPartition ++
		if sizeOfCurrentPartition >= partitionLimit { //Increment partition # if partition reaches max capacity while adding elements
			dirtyPartition = true
			partitionNumber++
			sizeOfCurrentPartition = 0
		}
	}

	// Set the partition number if we dirtied the partition number
	if dirtyPartition {
		client.Set(routingKey, partitionNumber, 0)
	}
	resp := &broker.ProduceResponse{}
	return resp, nil
}

func createRedisClient(addr string, password string, db int) (*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}