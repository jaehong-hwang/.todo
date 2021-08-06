package response

import (
	"fmt"
	"strings"

	"github.com/jaehong-hwang/todo/todo"
	"github.com/ryanuber/columnize"
)

// ListResponse is todo list response to string
type ListResponse struct {
	Todos todo.Todos
}

// Print todos by string like table
func (r *ListResponse) Print() {
	var output []string

	fields := todo.GetFields()
	output = append(output, strings.Join(fields[:], " | "))

	for _, todo := range r.Todos {
		fieldText := todo.ToStringSlice()
		output = append(output, strings.Join(fieldText[:], " | "))
	}

	fmt.Println(columnize.SimpleFormat(output))
}
