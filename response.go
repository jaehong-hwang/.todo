package main

import "fmt"

// Response interface
type Response interface {
	Print()
}

// MessageResponse struct
type MessageResponse struct {
	message string
}

// Print message
func (r *MessageResponse) Print() {
	fmt.Println(r.message)
}
