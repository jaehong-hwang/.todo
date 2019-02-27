package main

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"

	"github.com/iancoleman/strcase"
)

// TodoFileName is name of todo collection file
const TodoFileName string = ".todo"

// App is command center
type App struct {
	collection *TodoCollection
}

// NewApp func initial app
func NewApp() (App, error) {
	todoFile, _ := getTodoFile()

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

// getTodoDir return current directory has todo directory
func getTodoFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		_, err := os.Stat(dir + "/" + TodoFileName)
		if !os.IsNotExist(err) {
			return dir + "/" + TodoFileName, nil
		}

		dir = filepath.Dir(dir)
		if dir == "/" {
			return "", errors.New("todo collection doesn't exists, please run 'todo init'")
		}
	}
}
