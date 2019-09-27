package main

import (
	//"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
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
		for i := 0; i<5; i++ {
			routingKey = strings.TrimSuffix(routingKey, "\n")
			message := "example message set"
			_, err := Produce(routingKey, message+strconv.Itoa(i))
			if err != nil {
				panic(err)
			}
		}
		
	} else {
		for {
			err = Subscribe(routingKey)
			if err != nil {
				panic(err)
			}
		}
	}
}