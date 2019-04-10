package main

import (
	"errors"
	"os"
	"strings"
)

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	file  TodoFile
	todos []Todo
}

// Init todo collection directory
func (t *TodoCollection) Init() error {
	if t.file.IsExists() {
		return errors.New("todo collection already exists")
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = t.file.CreateFile(dir)
	if err != nil {
		return err
	}

	return nil
}

// Add todo item
func (t *TodoCollection) Add() error {
	input, err := t.file.GetContent()
	if err != nil {
		return err
	}

	Todo{
		id: t.getId()
	}

	lines := strings.Split(input, "\n")
	output := strings.Join(lines, "\n")

	err = t.file.FillContent(output)
	if err != nil {
		return err
	}

	return nil
}
