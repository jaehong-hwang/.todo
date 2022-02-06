package todo

import (
	"encoding/json"
	"strconv"

	"github.com/jaehong-hwang/todo/errors"
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
	id := 0
	if len(t.Todos) > 0 {
		id = t.Todos[len(t.Todos)-1].ID + 1
	}

	todo := Todo{
		ID:     id,
		Status: StatusWaiting,
		Level:  TodoLevels[0],
	}

	return todo
}

// Add todo to current collection
func (t *Collection) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

// Remove todo item by id
func (t *Collection) Remove(id int) bool {
	for i, todo := range t.Todos {
		if todo.ID == id {
			for j := i; j < len(t.Todos)-1; j++ {
				t.Todos[j] = t.Todos[j+1]
			}
			t.Todos = t.Todos[:len(t.Todos)-1]
			return true
		}
	}

	return false
}

// GetTodo by id
func (t *Collection) GetTodo(id int) (*Todo, error) {
	for i, todo := range t.Todos {
		if todo.ID == id {
			return &t.Todos[i], nil
		}
	}

	return nil, errors.NewWithParam("todo_id_not_found", map[string]string{
		"id": strconv.Itoa(id),
	})
}

// GetTodosJSONString from current collection
func (t *Collection) GetTodosJSONString() (string, error) {
	b, err := json.Marshal(t.Todos)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
