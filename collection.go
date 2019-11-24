package main

import (
	"encoding/json"
)

// Todos is todo array
type Todos []Todo

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	file  *File
	Todos Todos

	Args []string
}

// NewTodoCollection returned
func NewTodoCollection(todoFile *File) *TodoCollection {
	input, err := todoFile.GetContent()
	todos := Todos{}

	if err == nil {
		json.Unmarshal([]byte(input), &todos)
	}

	return &TodoCollection{
		file:  todoFile,
		Todos: todos,
	}
}

// Add todo item
func (t *TodoCollection) Add() {
	t.Todos = append(t.Todos, Todo{
		ID:      len(t.Todos),
		Content: t.Args[0],
	})

	if err := t.save(); err != nil {
		panic(err)
	}

	ResponseChan <- &MessageResponse{message: "add complete"}
}

// save todo items
func (t *TodoCollection) save() error {
	b, err := json.Marshal(t.Todos)
	if err != nil {
		return err
	}

	return t.file.FillContent(string(b))
}
