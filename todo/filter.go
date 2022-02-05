package todo

import (
	"strings"
)

// Filters is list filter of collection
type Filters struct {
	WithDone bool
	Status   []string
	Author   string
}

// GetList with current filter
func (f *Filters) GetList(t Todos) Todos {
	todos := Todos{}
	for _, todo := range t {
		isValidate := true

		if f.WithDone == false {
			isValidate = todo.Status == StatusWaiting || todo.Status == StatusWorking
		}

		if len(f.Status) > 0 {
			isValidate = false
			for _, s := range f.Status {
				if s == todo.Status {
					isValidate = true
					break
				}
			}
		}

		if f.Author != "" {
			isValidate = isValidate &&
				(strings.Contains(todo.Author, f.Author) ||
					strings.Contains(todo.AuthorEmail, f.Author))
		}

		if isValidate {
			todos = append(todos, todo)
		}
	}
	return todos
}
