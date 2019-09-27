package main

import (
	"google.golang.org/grpc"
)

func createGrpcClientConnection(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}