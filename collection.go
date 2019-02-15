package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

// TodoDirName is name of todo collection directory
const TodoDirName string = ".todo"

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	Dir string
}

// InitTodoCollection is make todo directory
func InitTodoCollection() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Mkdir(dir+"/"+TodoDirName, 0755)
	if err != nil {
		return err
	}

	return nil
}

// NewTodoCollection is TodoCollection's initialize function
func NewTodoCollection() (TodoCollection, error) {
	todo := TodoCollection{}

	dir, err := os.Getwd()
	if err != nil {
		return todo, err
	}

	todoDir, err := getTodoDir(dir)
	if err != nil {
		return todo, err
	}

	if todoDir == dir+"/"+TodoDirName {
		return todo, errors.New("todo collection already exists")
	}

	todo.Dir = dir

	return todo, err
}

// GetTodoDir return current directory has todo directory
func getTodoDir(dir string) (string, error) {
	for {
		_, err := os.Stat(dir + "/.todo")
		if !os.IsNotExist(err) {
			log.Println("todo dir: ", dir+"/"+TodoDirName)
			return dir, nil
		}

		dir = filepath.Dir(dir)
		if dir == "/" {
			return dir, errors.New("todo collection doesn't exists, please run 'todo init'")
		}
	}
}
