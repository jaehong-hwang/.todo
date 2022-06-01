package todo

import (
	"encoding/json"
	"sort"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/file"
	"github.com/rs/xid"
)

// Todos is todo array
type Todos []Todo

// Collection is manage .todo filesystem
type Collection struct {
	Todos Todos
}

// NewTodoCollection returned
func NewTodoCollection(todoFile *file.TodoWorkspace) *Collection {
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
		ID:     xid.New().String(),
		Status: STATUS_WAITING,
		Level:  TodoLevels[0],
	}

	return todo
}

// Add todo to current collection
func (t *Collection) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

// Remove todo item by id
func (t *Collection) Remove(id string) bool {
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
func (t *Collection) GetTodo(id string) (*Todo, error) {
	for i, todo := range t.Todos {
		if todo.ID == id {
			return &t.Todos[i], nil
		}
	}

	return nil, errors.NewWithParam("todo_id_not_found", map[string]string{
		"id": id,
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

// Sort collectio
func (t *Collection) Sort(orderBy string) error {
	if orderBy != "level" && orderBy != "due-date" && orderBy != "regist-date" {
		return errors.NewWithParam("sort_method_invalid", map[string]string{
			"sort": orderBy,
		})
	}

	sort.Slice(t.Todos, func(i, j int) bool {
		switch orderBy {
		case "level":
			return t.Todos[i].Level > t.Todos[j].Level
		case "due-date":
			return (t.Todos[i].DueDate.Unix() > 0 && t.Todos[j].DueDate.Unix() < 0) || t.Todos[i].DueDate.Unix() < t.Todos[j].DueDate.Unix()
		default:
			return t.Todos[i].RegistDate.Unix() < t.Todos[j].RegistDate.Unix()
		}
	})

	return nil
}
