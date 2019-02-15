package main

import (
	"fmt"
	"os"
)

func main() {
	// get command
	command := os.Args[1]

	// run command
	collection := TodoCollection{}
	err := collection.Run(command, os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
	}
}
