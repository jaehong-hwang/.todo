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

// Run comand
func (t *TodoCollection) Run(command string, args []string) error {
	if command == "init" {
		return t.Init()
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	todoDir, err := t.getTodoDir(dir)
	if err != nil {
		return err
	}

	if todoDir == dir+"/"+TodoDirName {
		return errors.New("todo collection already exists")
	}

	t.Dir = dir

	return nil
}

// Init todo collection directory
func (t *TodoCollection) Init() error {
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

// GetTodoDir return current directory has todo directory
func (t *TodoCollection) getTodoDir(dir string) (string, error) {
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
