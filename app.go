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
func NewApp() (*App, error) {
	todoFile, err := GetTodoFile()
	if err != nil {
		return nil, err
	}

	app := &App{
		collection: NewTodoCollection(todoFile),
	}

	return app, nil
}

// Run comand
func (a *App) Run(command string, args []string) (string, error) {
	command = strcase.ToCamel(command)

	_, ok := reflect.TypeOf(a.collection).MethodByName(command)
	if !ok {
		help, _ := a.collection.Help()
		return help, errors.New(command + " is invalid command")
	}

	a.collection.Args = args[1:]

	method := reflect.ValueOf(a.collection).MethodByName(command)
	result := method.Call([]reflect.Value{})

	if err, ok := result[1].Interface().(error); ok && err != nil {
		return "", err
	} else if response, ok := result[0].Interface().(string); ok {
		return response, nil
	}

	return "", errors.New("no receive responce")
}
