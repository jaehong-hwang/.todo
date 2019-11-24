package main

import (
	"os"
	"sync"
)

// ResponseChan for main app
var ResponseChan chan Response

func main() {
	var wg sync.WaitGroup

	// set ResponseChan
	ResponseChan = make(chan Response)

	wg.Add(1)

	app := NewApp()
	go app.Run(os.Args, &wg)

	go func() {
		wg.Wait()
		close(ResponseChan)
	}()

	response := <-ResponseChan
	if response != nil {
		response.Print()
	}
}
