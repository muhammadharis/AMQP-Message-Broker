package main

import (
	//"bufio"
	"fmt"
	"strings"
	"os"
)

func main() {
	
	cmdLineArguments := os.Args[1:]
	//fmt.Print("Enter routing key: ")
	//routingKey, err := reader.ReadString('\n')
	routingKey, err := "abc", error(nil)

	if cmdLineArguments[0] == "0" {
	//reader := bufio.NewReader(os.Stdin)
		if err != nil {
			panic(err)
		}
		routingKey = strings.TrimSuffix(routingKey, "\n")
		
		produceResponse, err := Produce(routingKey)
		if err != nil {
			panic(err)
		}
		fmt.Println(produceResponse)
	} else {
		err = Subscribe(routingKey)
		if err != nil {
			panic(err)
		}
	}
}