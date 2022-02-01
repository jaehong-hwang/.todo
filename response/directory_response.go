package response

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaehong-hwang/todo/todo"
)

// ListResponse is todo list response to string
type DirectoryResponse struct {
	Directories todo.Directories
}

// Print todos by string like table
func (r *DirectoryResponse) Print(isJson bool) {
	if isJson {
		b, err := json.Marshal(r.Directories)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println(string(b))
	} else {
		output := ""

		for _, directory := range r.Directories {
			output = output + directory + "\n"
		}

		fmt.Println(output)
	}
}
