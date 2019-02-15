package main

import (
	"fmt"
	"os"
)

func main() {
	// check command
	switch command := os.Args[1]; command {
	case "init":
		err := InitTodoCollection()
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println("todo collection intialized!")
		}
	default:
		// make todo collection
		collection, err := NewTodoCollection()
		if err != nil {
			fmt.Printf(err.Error())
		}

		fmt.Println(collection.Dir)
	}
}
