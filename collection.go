package main

import (
	"encoding/json"
	"errors"
	"os"
)

// Todos is todo array
type Todos []Todo

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	file  TodoFile
	todos Todos

	Args []string
}

// NewTodoCollection returned
func NewTodoCollection(todoFile TodoFile) (*TodoCollection, error) {
	input, err := todoFile.GetContent()
	todos := Todos{}

	if err == nil {
		json.Unmarshal([]byte(input), &todos)
	}

	t := &TodoCollection{
		file:  todoFile,
		todos: todos,
	}

	return t, nil
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

// Help command is show description for using todo app
func (t *TodoCollection) Help() (string, error) {
	return `usage: todo [--version] <command> [<args>]

Todo app helper.
You can run the following commands.

todo init		initial todo collection
todo add ${message}	adding todo`, nil
}

// Add todo item
func (t *TodoCollection) Add() (string, error) {
	t.todos = append(t.todos, Todo{
		ID:      len(t.todos),
		Content: t.Args[0],
	})

	if err := t.Save(); err != nil {
		return "", err
	}

	return "add complete", nil
}

// Save todo items
func (t *TodoCollection) Save() error {
	b, err := json.Marshal(t.todos)
	if err != nil {
		return err
	}

	return t.file.FillContent(string(b))
}
