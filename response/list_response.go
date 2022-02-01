package response

import (
	"fmt"
	"strings"

	"github.com/jaehong-hwang/todo/todo"
	"github.com/ryanuber/columnize"
)

// ListResponse is todo list response to string
type ListResponse struct {
	Collection todo.Collection
}

// Print todos by string like table
func (r *ListResponse) Print(isJson bool) {
	if isJson {
		str, _ := r.Collection.GetTodosJSONString()
		fmt.Println(str)
	} else {
		var output []string

		fields := todo.GetFields()
		output = append(output, strings.Join(fields[:], " | "))

		for _, todo := range r.Collection.Todos {
			fieldText := todo.ToStringSlice()
			output = append(output, strings.Join(fieldText[:], " | "))
		}

		fmt.Println(columnize.SimpleFormat(output))
	}
}
