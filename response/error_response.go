package response

import "fmt"

// ErrorResponse struct
type ErrorResponse struct {
	Err error
}

// Print error with ERROR tag
func (r *ErrorResponse) Print() {
	fmt.Println("[ERROR]", r.Err)
}
