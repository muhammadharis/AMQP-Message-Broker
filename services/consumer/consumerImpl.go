package consumer

import (
	"context"
	//"strconv"
	consumer "github.com/muhammadharis/grpc/protos/consumer"
	redis "github.com/go-redis/redis"
	helpers "github.com/muhammadharis/grpc/services/helpers"
)

// ConsumerImpl implements the server code for the Consumer proto
type ConsumerImpl struct {}

// ReadMessage allows sole consumers to read messages in a blocking fashion
func (*ConsumerImpl) ReadMessage(request *consumer.ConsumerReadRequest, stream consumer.ConsumerAPI_ReadMessageServer) error {
	queueName := request.QueueName
	client := helpers.CreateRedisClient("localhost:6379", "", 0) //No password, default DB

	xReadArguments := &redis.XReadArgs {
		Streams: []string{helpers.StreamNamespace+queueName, "$"},
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
					stream.Send(&consumer.ConsumerReadResponse{Message: s})
				}
			}
			
		}
	}

	return nil
}

func (*ConsumerImpl) CreateConsumerGroup(ctx context.Context, request *consumer.ConsumerGroupRequest) (*consumer.ConsumerGroupResponse, error) {
	return nil, nil
}
