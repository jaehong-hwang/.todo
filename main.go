package main

import (
	"log"
	"os"
)

func main() {
	// get command argument
	command := os.Args[1]

	// check command
	log.Println(command)

	// make todo collection
	collection := TodoCollection{}

	// check is todo directory
	if command != "init" && !collection.HasTodoDir() {
		log.Fatalf("You should run init first")
	}
}
