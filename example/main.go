package main

import (
	"fmt"
	"os"

	//broker "github.com/muhammadharis/grpc/protos/broker"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "publish" {
		fmt.Println("0")
		Message()
	} else if argsWithoutProg[0] == "subscribe" {
		fmt.Println("1")
		Read()
	}
}