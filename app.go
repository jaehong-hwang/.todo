package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sync"

	"github.com/iancoleman/strcase"
)

const (
	// TodoFileName is name of todo collection file
	todoFileName string = ".todo"

	// TodoFilePermission set read permission
	todoFilePermission os.FileMode = 0644

	// TodoNotFound error message
	todoNotFound string = "todo collection doesn't exists, please run 'todo init'"
)

// App is command center
type App struct {
	collection *TodoCollection
}

// RunCommand to running correct command
func RunCommand(command string, args []string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()

		if err := recover(); err != nil {
			fmt.Println(err)
			runtime.Goexit()
		}
	}()

	file := &File{name: ".todo", permission: 0644}
	if err := file.FindFromCurrentDirectory(); err != nil {
		panic(err)
	}

	a := &App{
		collection: NewTodoCollection(file),
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
