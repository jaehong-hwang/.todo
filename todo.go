package main

import (
	"strconv"
	"time"
)

// Todo unit struct
type Todo struct {
	id      int
	status  string
	author  string
	content string
	start   time.Time
	end     time.Time
}

const (
	// TodoAuthor string
	TodoAuthor = "Author"

	// TodoStart string
	TodoStart = "Start"

	// TodoEnd string
	TodoEnd = "End"

	// TodoContent string
	TodoContent = "Content"
)

// ToString from todo struct
func (t *Todo) ToString() string {
	todoStr := strconv.Itoa(t.id) + ". " + t.status + "\n"
	todoStr += TodoAuthor + " " + t.author + "\n"

	if !t.start.IsZero() {
		todoStr += TodoStart + " " + t.start.Format("yyyy-MM-dd") + "\n"
	}

	if !t.end.IsZero() {
		todoStr += TodoEnd + " " + t.end.Format("yyyy-MM-dd") + "\n"
	}

	todoStr += TodoContent + " " + t.content

	return todoStr
}
