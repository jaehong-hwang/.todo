package main

import (
	"os"

	"github.com/jaehong-hwang/todo/cli"
)

func main() {
	app := cli.NewApp()

	if response := app.Run(os.Args); response != nil {
		response.Print()
	}
}
