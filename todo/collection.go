package todo

import (
	"encoding/json"

	"github.com/jaehong-hwang/todo/file"
)

// Todos is todo array
type Todos []Todo

// Collection is manage .todo filesystem
type Collection struct {
	Todos Todos
}

// NewTodoCollection returned
func NewTodoCollection(todoFile *file.File) *Collection {
	todos := Todos{}

	if todoFile != nil {
		input, err := todoFile.GetContent()
		if err == nil {
			json.Unmarshal([]byte(input), &todos)
		}
	}

	return &Collection{
		Todos: todos,
	}
}

// NewTodo from todo list
func (t *Collection) NewTodo() Todo {
	todo := Todo{
		ID:     len(t.Todos),
		Status: StatusWaiting,
	}

	return todo
}

// Add todo to current collection
func (t *Collection) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

// Remove todo item by id
func (t *Collection) Remove(id int) {
	t.Todos[id] = t.Todos[len(t.Todos)-1]
	t.Todos = t.Todos[:len(t.Todos)-1]
}

// GetTodosByJSONString from current collection
func (t *Collection) GetTodosByJSONString() (string, error) {
	b, err := json.Marshal(t.Todos)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// GetTodosByStatus
func (t *Collection) GetTodosByStatus(status []string) Todos {
	todos := Todos{}
	for _, todo := range t.Todos {
		for _, s := range status {
			if s == todo.Status {
				todos = append(todos, todo)
				break
			}
		}
	}
	return todos
}
