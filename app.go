package main

import (
	"errors"
	"reflect"

	"github.com/iancoleman/strcase"
)

// App is command center
type App struct {
	collection *TodoCollection
}

// NewApp func initial app
func NewApp() (App, error) {
	todoFile, _ := GetTodoFile()

	collection := &TodoCollection{
		file: todoFile,
	}

	app := App{
		collection: collection,
	}

	return app, nil
}

// Run comand
func (a *App) Run(command string, args []string) error {
	command = strcase.ToCamel(command)

	_, ok := reflect.TypeOf(a.collection).MethodByName(command)
	if !ok {
		return errors.New(command + " is invalid command")
	}

	method := reflect.ValueOf(a.collection).MethodByName(command)
	method.Call([]reflect.Value{})

	return nil
}
