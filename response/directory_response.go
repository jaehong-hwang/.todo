package response

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaehong-hwang/todo/todo"
	"github.com/ryanuber/columnize"
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
		var output []string

		output = append(output, "id | name | path")

		for _, directory := range r.Directories {
			output = append(output, directory.ID + " | " + directory.Name + " | " + directory.Path)
		}

		fmt.Println(columnize.SimpleFormat(output))
	}
}
