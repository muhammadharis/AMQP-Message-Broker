package main

import (
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "publish" {
		Message()
	} else if argsWithoutProg[0] == "subscribe" {
		Read()
	}
}