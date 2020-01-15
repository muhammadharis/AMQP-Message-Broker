FROM golang:1.13 AS builder

WORKDIR $GOPATH/src/github.com/muhammadharis/AMQP-Message-Broker

COPY . ./
RUN go get -d ./...
RUN go install ./...

CMD ["/go/bin/AMQP-Message-Broker"]
