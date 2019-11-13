package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ryanuber/columnize"
)

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

// ListResponse is todo list response to string
type ListResponse struct {
	todos Todos
}

// Print todos by string like table
func (r *ListResponse) Print() {
	var fields []string
	var output []string

	val := reflect.Indirect(reflect.ValueOf(Todo{}))
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}

	output = append(output, strings.Join(fields[:], " | "))
	for _, todo := range r.todos {
		var fieldText []string
		for _, field := range fields {
			str := fmt.Sprintf("%v", reflect.Indirect(reflect.ValueOf(todo)).FieldByName(field).Interface())
			fieldText = append(fieldText, str)
		}
		output = append(output, strings.Join(fieldText[:], " | "))
	}

	fmt.Println(columnize.SimpleFormat(output))
}
