package main

import (
	"os"

	"github.com/jaehong-hwang/todo/cli"
)

func main() {
	if response, isJson := cli.Run(os.Args); response != nil {
		response.Print(isJson)
	}
}
