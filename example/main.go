package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "0" {
		fmt.Println("0")
		Message()
	} else if argsWithoutProg[0] == "1" {
		fmt.Println("1")
		Read()
	}
}