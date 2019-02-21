package main

import (
	"fmt"
	"os"
)

func main() {
	// get command
	command := os.Args[1]

	// run command
	app, err := NewApp()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = app.Run(command, os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
	}
}
