# Message Broker Service Implemented in gRPC

This is my implementation of AMQP done in Go using gRPC. The purpose of this project was to understand gRPC and how I could integrate Redis streaming with gRPC server-streaming.

[broker.proto](/protos/broker/broker.proto) describes the two services (`ProduceAPI` and `ConsumerAPI`), alongside any necessary information needed to create requests.

[The example folder](/example/) contains examples for what the client-implementation of Producer or Consumer might look like.
