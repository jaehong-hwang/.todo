package main

import (
	"os"

	"github.com/jaehong-hwang/todo/cli"
	r "github.com/jaehong-hwang/todo/response"
)

func main() {
	if response, isJson := cli.Run(os.Args); response != nil || isJson {
		if isJson && response == nil {
			response = &r.MessageResponse{Message: "empty response"}
		}
		response.Print(isJson)
	}
}
