FROM golang:1.13 AS builder

WORKDIR $GOPATH/src/github.com/muhammadharis/AMQP-Message-Broker/

COPY . ./
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./

ENTRYPOINT ["./app"]