package response

import (
	"fmt"
	"os"
)

// ErrorResponse struct
type ErrorResponse struct {
	Err error
}

// Print error with ERROR tag
func (r *ErrorResponse) Print() {
	fmt.Fprintln(os.Stderr, r.Err)
	os.Exit(1)
}
