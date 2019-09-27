package main

import (
	//"bufio"
	"fmt"
	"strings"
	//"os"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter routing key: ")
	//routingKey, err := reader.ReadString('\n')
	routingKey, err := "abc", error(nil)
	if err != nil {
		panic(err)
	}
	routingKey = strings.TrimSuffix(routingKey, "\n")
	
	produceResponse, err := Produce(routingKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(produceResponse)
}