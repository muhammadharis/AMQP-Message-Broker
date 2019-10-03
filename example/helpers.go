package main

import (
	"google.golang.org/grpc"
)

// CreateGrpcClientConnection creates a client connection to the gRPC server
func CreateGrpcClientConnection(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}