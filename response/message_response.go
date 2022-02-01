package response

import (
	"encoding/json"
	"fmt"
)

// MessageResponse struct
type MessageResponse struct {
	Message string
}

// Print message
func (r *MessageResponse) Print(isJson bool) {
	if isJson {
		str, _ := json.Marshal(r)
		fmt.Println(string(str))
	} else {
		fmt.Println(r.Message)
	}
}
