package main

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"

	"github.com/iancoleman/strcase"
)

// TodoDirName is name of todo collection directory
const TodoDirName string = ".todo"

// App is command center
type App struct {
	collection *TodoCollection
}

// NewApp func initial app
func NewApp() (App, error) {
	todoDir, err := getTodoDir()
	if err != nil {
		return App{}, err
	}

	collection := &TodoCollection{
		dir: todoDir,
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
func getTodoDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		_, err := os.Stat(dir + "/" + TodoDirName)
		if !os.IsNotExist(err) {
			return dir + "/" + TodoDirName, nil
		}

		dir = filepath.Dir(dir)
		if dir == "/" {
			return "", errors.New("todo collection doesn't exists, please run 'todo init'")
		}
	}
}
