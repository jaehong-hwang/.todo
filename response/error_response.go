package response

import (
	"log"
)

// ErrorResponse struct
type ErrorResponse struct {
	Err error
}

// Print error with ERROR tag
func (r *ErrorResponse) Print() {
	log.Fatalf("[ERROR] %s", r.Err)
}
