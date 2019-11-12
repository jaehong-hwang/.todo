package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/iancoleman/strcase"
)

// App is command center
type App struct {
	collection *TodoCollection
}

// RunCommand to running correct command
func RunCommand(command string, args []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			runtime.Goexit()
		}
	}()

	todoFile, err := GetTodoFile()
	if err != nil {
		panic(err)
	}

	a := &App{
		collection: NewTodoCollection(todoFile),
	}

	command = strcase.ToCamel(command)

	_, ok := reflect.TypeOf(a.collection).MethodByName(command)
	if !ok {
		a.collection.Help()
	}

	a.collection.Args = args[1:]

	method := reflect.ValueOf(a.collection).MethodByName(command)
	method.Call([]reflect.Value{})
}
