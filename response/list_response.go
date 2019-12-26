package response

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jaehong-hwang/todo/todo"
	"github.com/ryanuber/columnize"
)

// ListResponse is todo list response to string
type ListResponse struct {
	Todos todo.Todos
}

// Print todos by string like table
func (r *ListResponse) Print() {
	var fields []string
	var output []string

	val := reflect.Indirect(reflect.ValueOf(todo.Todo{}))
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}

	output = append(output, strings.Join(fields[:], " | "))
	for _, todo := range r.Todos {
		var fieldText []string
		for _, field := range fields {
			str := fmt.Sprintf("%v", reflect.Indirect(reflect.ValueOf(todo)).FieldByName(field).Interface())
			fieldText = append(fieldText, str)
		}
		output = append(output, strings.Join(fieldText[:], " | "))
	}

	fmt.Println(columnize.SimpleFormat(output))
}