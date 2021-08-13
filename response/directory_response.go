package response

import (
	"fmt"

	"github.com/jaehong-hwang/todo/todo"
)

// ListResponse is todo list response to string
type DirectoryResponse struct {
	Directories todo.Directories
}

// Print todos by string like table
func (r *DirectoryResponse) Print() {
	output := ""

	for _, directory := range r.Directories {
		output = output + directory + "\n"
	}

	fmt.Println(output)
}
