package todo

import (
	"reflect"
	"strconv"
	"time"
)

const (
	StatusWaiting = "waiting"
	StatusWorking = "working"
	StatusDone    = "done"
)

// Todo unit struct
type Todo struct {
	ID          int       `json:"id"`
	Status      string    `json:"status"`
	Author      string    `json:"author"`
	AuthorEmail string    `json:"authorEmail"`
	Content     string    `json:"content"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

func GetFields() []string {
	var fields []string
	val := reflect.Indirect(reflect.ValueOf(Todo{}))
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}

	return fields
}

func (t *Todo) ToStringSlice() []string {
	return []string{
		strconv.Itoa(t.ID),
		t.Status,
		t.Author,
		t.AuthorEmail,
		t.Content,
		t.Start.Format("2006.01.02 15:04"),
		t.End.Format("2006.01.02 15:04"),
	}
}
