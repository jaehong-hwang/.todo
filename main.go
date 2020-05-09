package main

import (
	"os"

	"github.com/jaehong-hwang/todo/cli"
)

func main() {
	if response := cli.Run(os.Args); response != nil {
		response.Print()
	}
}
