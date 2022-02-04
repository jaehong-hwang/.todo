package main

import (
	"github.com/jaehong-hwang/todo/cmd"
	r "github.com/jaehong-hwang/todo/response"
)

func main() {
	if response, isJson := cmd.Execute(); response != nil || isJson {
		if isJson && response == nil {
			response = &r.MessageResponse{Message: "empty response"}
		}
		response.Print(isJson)
	}
}
