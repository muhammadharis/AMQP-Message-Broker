package broker

import (
	"fmt"
	"context"
	//"strconv"
	broker "github.com/muhammadharis/grpc/protos/broker"
	redis "github.com/go-redis/redis"
	
	helpers "github.com/muhammadharis/grpc/services/helpers"
)

const partitionLimit = 10

type BrokerImpl struct {}

// Produce consumes a producer's message
func (*BrokerImpl) Produce(ctx context.Context, request *broker.ProduceRequest) (*broker.ProduceResponse, error) {
	routingKey := request.RoutingKey
	//messageSet := request.MessageSet
	client := helpers.CreateRedisClient("localhost:6379", "", 0) //No password, default DB
	
	xAddArgument := &redis.XAddArgs{
		Stream: routingKey,
		MaxLen: partitionLimit,
		MaxLenApprox: partitionLimit,
		ID: "*",
		Values: map[string]interface{} {
			"a" : 1,
			"b" : 2,
			"c" : 3,
			"d" : 4,
			"e" : 5,
			"f" : 6,
			"g" : 7,
			"h" : 2,
			"i" : 1,
			"j" : 2,
			"k" : 1,
			"l" : 2,
			"m" : 1,
			"n" : 2,
			"o" : 1,
			"p" : 2,
		},
	}
	cmd := client.XAdd(xAddArgument)
	fmt.Println(cmd.String())

	/*dirtyPartition := false // Is set when the partition number is changed

	// The Redis routing key will hold the partition number, actual data is stored in routingkey+partitionNumber
	var partitionNumber int64
	if partitionUnparsed := client.Get(routingKey); partitionUnparsed.Val() == "" {
		partitionNumber = 0
	} else {
		var err error
		partitionNumber, err = partitionUnparsed.Int64()
		if err != nil {
			panic(err)
		}
	}

	sizeOfCurrentPartition := client.LLen(routingKey+string(partitionNumber)).Val()
	// Increment the partition number if the partition is at its max capacity
	if sizeOfCurrentPartition >= partitionLimit {
		dirtyPartition = true
		partitionNumber++
		sizeOfCurrentPartition = 0
	}

	for msg := range messageSet {
		client.RPush(routingKey+strconv.FormatInt(partitionNumber, 10), msg) //We store messages in routingKey+partitionNumber
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
	}*/
	resp := &broker.ProduceResponse{}
	return resp, nil
}

