package main

import (
	"errors"
	"os"
)

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	dir string
}

// Init todo collection directory
func (t *TodoCollection) Init() error {
	if t.dir != "" {
		return errors.New("todo collection already exists")
	}

	err := os.Mkdir(t.dir, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Add todo item
func (t *TodoCollection) Add() error {
	f, err := os.OpenFile(t.dir+"/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.WriteString("text")

	return nil
}
