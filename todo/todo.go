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
	Labels      Labels    `json:"label"`
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
		t.Labels.ToString(),
	}
}

func (t *Todo) AddLabel(l *Label) bool {
	for _, lb := range t.Labels {
		if lb.Text == l.Text {
			return false
		}
	}

	t.Labels = append(t.Labels, l)
	return true
}

func (t *Todo) RemoveLabel(labelText string) bool {
	for i, l := range t.Labels {
		if l.Text == labelText {
			for j := i; j < len(t.Labels)-1; j++ {
				t.Labels[j] = t.Labels[j+1]
			}
			t.Labels = t.Labels[:len(t.Labels)-1]
			return true
		}
	}

	return false
}
