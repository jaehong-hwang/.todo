package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// TodoFileName is name of todo collection file
	todoFileName string = ".todo"

	// TodoFilePermission set read permission
	todoFilePermission os.FileMode = 0644
)

// TodoFile have role todo file read, write
type TodoFile struct {
	path string
}

// IsExists todo file in current directory
func (t *TodoFile) IsExists() bool {
	return t.path != "/"
}

// GetContent from todo file
func (t *TodoFile) GetContent() (string, error) {
	content, err := ioutil.ReadFile(t.path)
	return string(content), err
}

// FillContent to todo file
func (t *TodoFile) FillContent(content string) error {
	return ioutil.WriteFile(t.path, []byte(content), todoFilePermission)
}

// CreateFile of tood
func (t *TodoFile) CreateFile(dir string) error {
	f, err := os.Create(dir + "/" + todoFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

// GetTodoFile return current directory has todo directory
func GetTodoFile() (TodoFile, error) {
	dir, err := os.Getwd()
	if err != nil {
		return TodoFile{}, err
	}

	for {
		_, err := os.Stat(dir + "/" + todoFileName)
		if !os.IsNotExist(err) {
			return TodoFile{
				path: dir + "/" + todoFileName,
			}, nil
		}

		dir = filepath.Dir(dir)
		if dir == "/" {
			return TodoFile{}, errors.New("todo collection doesn't exists, please run 'todo init'")
		}
	}
}
