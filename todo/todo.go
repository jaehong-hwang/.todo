package todo

import (
	"time"
)

const (
	StatusWaiting = "waiting"
	StatusWorking = "working"
	StatusDone    = "done"
)

// Todo unit struct
type Todo struct {
	ID      int       `json:"id"`
	Status  string    `json:"status"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
}
