# AMQP Message Broker Service Implemented in gRPC

This is my implementation of an AMQP Message Broker done in Go using gRPC. The purpose of this project was to understand gRPC and how I could integrate Redis streaming with gRPC server-streaming.

[broker.proto](/protos/broker/broker.proto) describes the two services (`ProduceAPI` and `ConsumerAPI`), alongside any necessary information needed to create requests.

[The example folder](/example/) contains examples for what the client-implementation of Producer or Consumer might look like.

The idea for this messaging service is similar to that of RabbitMQ. In this implementation, there are 2 types of exchanges: **Fanout** and **Direct**. The basic architecture is shown.
![AMQP Architecture](https://callistaenterprise.se/assets/blogg/goblog/part9-rabbitmq-exchange.png)

# Containerization
### Build the image:
docker build -t muhammadharis/amqp-message-broker .
```bash
docker build -t muhammadharis/amqp-message-broker .
```

### Run the image:
```bash
docker run -d --name messagebrokerapp muhammadharis/amqp-message-broker
```
