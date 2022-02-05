package todo

import (
	"reflect"
	"strconv"
	"time"

	"github.com/jaehong-hwang/todo/errors"
)

const (
	StatusWaiting = "waiting"
	StatusWorking = "working"
	StatusDone    = "done"
)

var (
	TodoLevels = []int{0, 1, 2, 3}
)

func IsValidStatus(status string) bool {
	return status == StatusWaiting || status == StatusWorking || status == StatusDone
}

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
	Level       int       `json:"level"`
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
		strconv.Itoa(t.Level),
	}
}

func (t *Todo) AddLabel(l *Label) error {
	for _, lb := range t.Labels {
		if lb.Text == l.Text {
			return errors.NewWithParam("label_already_exists", map[string]string{
				"label": l.Text,
			})
		}
	}

	t.Labels = append(t.Labels, l)
	return nil
}

func (t *Todo) RemoveLabel(labelText string) error {
	for i, l := range t.Labels {
		if l.Text == labelText {
			for j := i; j < len(t.Labels)-1; j++ {
				t.Labels[j] = t.Labels[j+1]
			}
			t.Labels = t.Labels[:len(t.Labels)-1]
			return nil
		}
	}

	return errors.NewWithParam("label_not_found", map[string]string{
		"label": labelText,
	})
}
