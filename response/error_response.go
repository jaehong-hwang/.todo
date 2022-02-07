package response

import (
	"encoding/json"
	"fmt"
	"os"
)

// ErrorResponse struct
type ErrorResponse struct {
	Err error
}

// Print error with ERROR tag
func (r *ErrorResponse) Print(isJson bool) {
	if isJson {
		str, _ := json.Marshal(r)
		fmt.Fprintln(os.Stderr, string(str))
	} else {
		fmt.Fprintln(os.Stderr, r.Err)
	}
	os.Exit(1)
}
