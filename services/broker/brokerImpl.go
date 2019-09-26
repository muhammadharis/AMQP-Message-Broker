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

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	dirtyPartition := false

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
