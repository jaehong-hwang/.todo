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
func (t *TodoCollection) Add(todo Todo) error {
	todo.ID = len(t.Todos)
	t.Todos = append(t.Todos, todo)

	if err := t.save(); err != nil {
		return err
	}

	return nil
}

// save todo items
func (t *TodoCollection) save() error {
	b, err := json.Marshal(t.Todos)
	if err != nil {
		return err
	}

	return t.file.FillContent(string(b))
}
