package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Todos is todo array
type Todos []Todo

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	file  TodoFile
	todos Todos

	Args []string
}

// NewTodoCollection returned
func NewTodoCollection(todoFile TodoFile) (*TodoCollection, error) {
	input, err := todoFile.GetContent()
	todos := Todos{}

	if err == nil {
		json.Unmarshal([]byte(input), &todos)
	}

	t := &TodoCollection{
		file:  todoFile,
		todos: todos,
	}

	return t, nil
}

// Init todo collection directory
func (t *TodoCollection) Init() (string, error) {
	if t.file.IsExists() {
		return "", errors.New("todo collection already exists")
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	err = t.file.CreateFile(dir)
	if err != nil {
		return "", err
	}

	return "todo init complete", nil
}

// Help command is show description for using todo app
func (t *TodoCollection) Help() (string, error) {
	return `usage: todo [--version] <command> [<args>]

Todo app helper.
You can run the following commands.

todo init		initial todo collection
todo add ${message}	adding todo`, nil
}

// List of todo items
func (t *TodoCollection) List() (string, error) {
	var fields []string

	val := reflect.Indirect(reflect.ValueOf(Todo{}))
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}

	fmt.Println(strings.Join(fields[:], "\t "))
	for _, todo := range t.todos {
		for _, field := range fields {
			fmt.Printf("%v\t", reflect.Indirect(reflect.ValueOf(todo)).FieldByName(field).Interface())

		}
		fmt.Printf("\n")
	}

	return "", nil
}

// Add todo item
func (t *TodoCollection) Add() (string, error) {
	t.todos = append(t.todos, Todo{
		ID:      len(t.todos),
		Content: t.Args[0],
	})

	if err := t.save(); err != nil {
		return "", err
	}

	return "add complete", nil
}

// save todo items
func (t *TodoCollection) save() error {
	b, err := json.Marshal(t.todos)
	if err != nil {
		return err
	}

	return t.file.FillContent(string(b))
}
