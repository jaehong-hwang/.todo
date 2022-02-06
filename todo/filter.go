package todo

import (
	"strings"
	"time"
)

// Filters is list filter of collection
type Filters struct {
	WithDone     bool
	Status       []string
	Author       string
	DueDateStart time.Time
	DueDateEnd   time.Time
}

// Run filter, return filtered collection
func (f *Filters) Run(c *Collection) *Collection {
	col := Collection{}
	for _, todo := range c.Todos {
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

		startYear, startMonth, startDate := f.DueDateStart.Date()
		isValidate = isValidate && todo.DueDate.Year() >= startYear && int(todo.DueDate.Month()) >= int(startMonth) && todo.DueDate.Day() >= startDate

		if f.DueDateEnd.Year() > 1 {
			endYear, endMonth, endDate := f.DueDateEnd.Date()
			isValidate = isValidate && todo.DueDate.Year() <= endYear && int(todo.DueDate.Month()) <= int(endMonth) && todo.DueDate.Day() <= endDate
		}

		if isValidate {
			col.Todos = append(col.Todos, todo)
		}
	}
	return &col
}
