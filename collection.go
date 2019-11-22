package main

import (
	"encoding/json"
	"os"
)

// Todos is todo array
type Todos []Todo

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	file  *File
	todos Todos

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
		todos: todos,
	}
}

// Init todo collection directory
func (t *TodoCollection) Init() {
	if t.file.IsExists() {
		panic("todo collection already exists")
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = t.file.CreateFile(dir)
	if err != nil {
		panic(err)
	}

	ResponseChan <- &MessageResponse{message: "todo init complete"}
}

// Help command is show description for using todo app
func (t *TodoCollection) Help() {
	ResponseChan <- &MessageResponse{message: `usage: todo [--version] <command> [<args>]

Todo app helper.
You can run the following commands.

todo init		initial todo collection
todo add ${message}	adding todo`}
}

// List of todo items
func (t *TodoCollection) List() {
	ResponseChan <- &ListResponse{todos: t.todos}
}

// Add todo item
func (t *TodoCollection) Add() {
	t.todos = append(t.todos, Todo{
		ID:      len(t.todos),
		Content: t.Args[0],
	})

	if err := t.save(); err != nil {
		panic(err)
	}

	ResponseChan <- &MessageResponse{message: "add complete"}
}

// save todo items
func (t *TodoCollection) save() error {
	b, err := json.Marshal(t.todos)
	if err != nil {
		return err
	}

	return t.file.FillContent(string(b))
}
