package consumer

import (
	//"strconv"
	"context"
	broker "github.com/muhammadharis/AMQP-Message-Broker/protos/broker"
	redis "github.com/go-redis/redis"
	helpers "github.com/muhammadharis/AMQP-Message-Broker/services/helpers"
)

// ConsumerImpl implements the server code for the Consumer proto
type ConsumerImpl struct {}

// ReadMessage allows sole consumers to read messages in a blocking fashion
func (*ConsumerImpl) ReadMessage(request *broker.ConsumerReadRequest, stream broker.ConsumerAPI_ReadMessageServer) error {
	queueName := request.QueueName
	client := helpers.CreateRedisClient("", 0) //No password, default DB

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
					stream.Send(&broker.ConsumerReadResponse{Message: s})
				}
			}
			
		}
	}

	return nil
}

// GroupReadMessage allows consumer groups read messages in a blocking fashion
func (*ConsumerImpl) GroupReadMessage(request *broker.ConsumerGroupReadRequest, stream broker.ConsumerAPI_GroupReadMessageServer) error {
	consumerGroupName := request.ConsumerGroupName
	consumerName := request.ConsumerName
	queueName := request.QueueName
	client := helpers.CreateRedisClient("", 0) //No password, default DB

	xReadGroupArguments := &redis.XReadGroupArgs {
		Group: consumerGroupName,
		Consumer: consumerName,
		Streams: []string{helpers.StreamNamespace+queueName, "$"},
		Block: 0,
	}

	xStreamSliceCmd := client.XReadGroup(xReadGroupArguments)
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
					stream.Send(&broker.ConsumerGroupReadResponse{Message: s})
				}
			}
			
		}
	}
	return nil
}

// CreateConsumerGroup creates a new consumer group on the server
func (*ConsumerImpl) CreateConsumerGroup(ctx context.Context, request *broker.CreateConsumerGroupRequest) (*broker.CreateConsumerGroupResponse, error) {
	groupName := request.GroupName
	queueName := request.QueueName
	client := helpers.CreateRedisClient("", 0) //No password, default DB

	client.XGroupCreate(helpers.StreamNamespace+queueName, groupName, "$")

	return &broker.CreateConsumerGroupResponse{}, nil
}
