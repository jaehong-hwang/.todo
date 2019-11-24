package todo

import (
	"encoding/json"

	"github.com/jaehong-hwang/todo/file"
)

// Todos is todo array
type Todos []Todo

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	Todos Todos
}

// NewTodoCollection returned
func NewTodoCollection(todoFile *file.File) *TodoCollection {
	input, err := todoFile.GetContent()
	todos := Todos{}

	if err == nil {
		json.Unmarshal([]byte(input), &todos)
	}

	return &TodoCollection{
		Todos: todos,
	}
}

// NewTodo from todo list
func (t *TodoCollection) NewTodo() Todo {
	todo := Todo{
		ID: len(t.Todos),
	}

	return todo
}

// Add todo to current collection
func (t *TodoCollection) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

// GetTodosByJSONString from current collection
func (t *TodoCollection) GetTodosByJSONString() (string, error) {
	b, err := json.Marshal(t.Todos)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
