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

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	dir string
}

// Run comand
func (t *TodoCollection) Run(command string, args []string) error {
	if command == "init" {
		return t.Init()
	}

	command = strcase.ToCamel(command)

	_, ok := reflect.TypeOf(t).MethodByName(command)
	if !ok {
		return errors.New(command + " is invalid command")
	}

	method := reflect.ValueOf(t).MethodByName(command)
	method.Call([]reflect.Value{})

	return nil
}

// Init todo collection directory
func (t *TodoCollection) Init() error {
	todoDir, err := t.getTodoDir()
	if err != nil {
		return err
	}

	if todoDir != "" {
		return errors.New("todo collection already exists")
	}

	err = os.Mkdir(todoDir, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Add todo item
func (t *TodoCollection) Add() error {
	dir, err := t.getTodoDir()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(dir+"/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.WriteString("text")

	return nil
}

// GetTodoDir return current directory has todo directory
func (t *TodoCollection) getTodoDir() (string, error) {
	if t.dir == "" {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}

		for {
			_, err := os.Stat(dir + "/" + TodoDirName)
			if !os.IsNotExist(err) {
				t.dir = dir + "/" + TodoDirName
				return t.dir, nil
			}

			dir = filepath.Dir(dir)
			if dir == "/" {
				return "", errors.New("todo collection doesn't exists, please run 'todo init'")
			}
		}
	} else {
		return t.dir, nil
	}
}
