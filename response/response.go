package response

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jaehong-hwang/todo/todo"
	"github.com/ryanuber/columnize"
)

// Response interface
type Response interface {
	Print()
}

// MessageResponse struct
type MessageResponse struct {
	Message string
}

// Print message
func (r *MessageResponse) Print() {
	fmt.Println(r.Message)
}

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

// ErrorResponse struct
type ErrorResponse struct {
	Err error
}

// Print error with ERROR tag
func (r *ErrorResponse) Print() {
	fmt.Println("[ERROR]", r.Err)
}
