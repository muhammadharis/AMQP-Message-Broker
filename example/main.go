package main

import (
	"os"

	cli "github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Simple pizza CLI"
	app.Usage = "An example CLI for ordering pizza's"
	app.Author = "muhammadharis" 
}

func main() {
	info()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}