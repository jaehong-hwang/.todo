package main

import (
	"log"
	"os"
	"path/filepath"
)

// TodoCollection is manage .todo filesystem
type TodoCollection struct{}

// HasTodoDir return current directory has todo directory
func (t *TodoCollection) HasTodoDir() bool {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err = os.Stat(dir + "/.todo")
		if !os.IsNotExist(err) {
			log.Println("todo dir: ", dir+"/.todo")
			return true
		}

		dir = filepath.Dir(dir)
		if dir == "/" {
			return false
		}
	}
}
