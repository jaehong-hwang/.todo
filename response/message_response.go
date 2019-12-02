package response

import "fmt"

// MessageResponse struct
type MessageResponse struct {
	Message string
}

// Print message
func (r *MessageResponse) Print() {
	fmt.Println(r.Message)
}
