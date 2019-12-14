package todo

import (
	"time"
)

const (
	StatusWaiting = 0
	StatusWorking = 1
	StatusDone    = 2
)

// Todo unit struct
type Todo struct {
	ID      int       `json:"id"`
	Status  int8      `json:"status"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
}
