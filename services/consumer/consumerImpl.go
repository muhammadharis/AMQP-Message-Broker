package consumer

import (
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
		Streams: []string{key, "$"},
		Block: 0,
	}
	
	xStreamSliceCmd := client.XRead(xReadArguments)
	xStreams, err := xStreamSliceCmd.Result() //Get Results from XRead command
	if err != nil {
		panic(err)
	}

	for _, xStream := range xStreams { //Get individual xStream
		//streamName := xStream.Stream
		for _, xMessage := range xStream.Messages { // Get the message from the xStream
			for _, v := range xMessage.Values { // Get the values from the message
				s, ok := v.(string)
				if ok {
					stream.Send(&consumer.ConsumerResponse{Message: s})
				}
			}
			
		}
	}

	return nil
}
