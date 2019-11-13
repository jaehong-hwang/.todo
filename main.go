package main

import (
	"os"
	"sync"
)

// ResponseChan for main app
var ResponseChan chan Response

func main() {
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "help")
	}

	var wg sync.WaitGroup

	// set ResponseChan
	ResponseChan = make(chan Response)

	// get command
	command := os.Args[1]

	wg.Add(1)
	go RunCommand(command, os.Args[1:], &wg)

	go func() {
		wg.Wait()
		close(ResponseChan)
	}()

	response := <-ResponseChan
	if response != nil {
		response.Print()
	}
}
