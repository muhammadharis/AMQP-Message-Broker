package consumer

import (
	"fmt"
	//"strconv"
	consumer "github.com/muhammadharis/grpc/protos/consumer"
	redis "github.com/go-redis/redis"
	helpers "github.com/muhammadharis/grpc/services/helpers"
)

type ConsumerImpl struct {}

// Subscribe allows producers to subscribe to messages
func (*ConsumerImpl) Subscribe(request *consumer.ConsumerRequest, stream consumer.ConsumerAPI_SubscribeServer) error {
	key := request.Key
	client := helpers.CreateRedisClient("localhost:6379", "", 0) //No password, default DB

	xReadArguments := &redis.XReadArgs {
		Streams: []string{key},
		Block: 0,
	}
	
	xStreamSlice := client.XRead(xReadArguments)
	fmt.Println(xStreamSlice.String())
	return nil
}
