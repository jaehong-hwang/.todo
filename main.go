package main

import (
	"os"
)

// ResponseChan for main app
var ResponseChan chan Response

func main() {
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "help")
	}

	// set ResponseChan
	ResponseChan = make(chan Response)

	// get command
	command := os.Args[1]

	go RunCommand(command, os.Args[1:])

	response := <-ResponseChan
	response.Print()
}
