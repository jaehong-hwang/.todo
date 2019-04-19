package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Hi, I'm Todo manager in your cli")
	} else {
		// get command
		command := os.Args[1]

		// run command
		app, err := NewApp()
		if err != nil {
			fmt.Println(err.Error())
		}

		res, err := app.Run(command, os.Args[1:])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}
	}
}
