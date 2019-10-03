# AMQP Message Broker Service Implemented in gRPC

This is my implementation of an AMQP Message Broker done in Go using gRPC. The purpose of this project was to understand gRPC and how I could integrate Redis streaming with gRPC server-streaming.

[broker.proto](/protos/broker/broker.proto) describes the two services (`ProduceAPI` and `ConsumerAPI`), alongside any necessary information needed to create requests.

[The example folder](/example/) contains examples for what the client-implementation of Producer or Consumer might look like.

The idea for this messaging service is similar to that of RabbitMQ. In this implementation, there are 2 types of exchanges: **Fanout** and **Direct**. The basic architecture is shown.
![AMQP Architecture](https://callistaenterprise.se/assets/blogg/goblog/part9-rabbitmq-exchange.png)

# Running an Example
Have three separate terminal instances. 
#### Step 1:
Flush the Redis cache by running `flushall` in the Redis CLI

#### Step 2:
In the first terminal, start the server, by running
```bash
go run main.go
``` 
in the root directory.

#### Step 3:
In the [example directory](/example/), run `make`. An executable named `example` will be generated. To test the service, run
```bash
./example subscribe
```
in one terminal, and 
```bash
./example publish
```
in the other terminal. 

#### Step 4:
You will notice that the `subscribe` will wait until the message is published, and then print. The client code can be modified for any use case.

#### Demo
![Broker Demo](https://user-images.githubusercontent.com/13709152/66133635-fd727c00-e5c4-11e9-92ea-46f62d348b0f.gif)

# Containerization
### Build the image:
docker build -t muhammadharis/amqp-message-broker .
```bash
docker build -t muhammadharis/amqp-message-broker .
```

### Run a container from the image:
```bash
docker run -d --name messagebrokerapp muhammadharis/amqp-message-broker
```
