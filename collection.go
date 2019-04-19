package main

import (
	"errors"
	"os"
)

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	file  TodoFile
	todos []Todo

	Args []string
}

// Init todo collection directory
func (t *TodoCollection) Init() (string, error) {
	if t.file.IsExists() {
		return "", errors.New("todo collection already exists")
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	err = t.file.CreateFile(dir)
	if err != nil {
		return "", err
	}

	return "todo init complete", nil
}

// Add todo item
func (t *TodoCollection) Add() (string, error) {
	input, err := t.file.GetContent()
	if err != nil {
		return "", err
	}

	/*t.todos = append(t.todos, Todo{
		id: len(t.todos),
	})*/

	output := input + t.Args[0] + "\n"

	err = t.file.FillContent(output)
	if err != nil {
		return "", err
	}

	return "add complete", nil
}
