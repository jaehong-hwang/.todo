package file

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

	// TodoNotFound error message
	todoNotFound string = "todo collection doesn't exists, please run 'todo init'"
)

// File management struct
type File struct {
	Name       string
	Permission os.FileMode
	path       string
}

// IsExists todo file in current directory
func (f *File) IsExists() bool {
	return f.path != "" && f.path != "/"
}

// GetContent from todo file
func (f *File) GetContent() (string, error) {
	if f.IsExists() == false {
		return "", errors.New(todoNotFound)
	}

	content, err := ioutil.ReadFile(f.path)
	return string(content), err
}

// FillContent to todo file
func (f *File) FillContent(content string) error {
	if f.IsExists() == false {
		return errors.New(todoNotFound)
	}

	return ioutil.WriteFile(f.path, []byte(content), f.Permission)
}

// CreateFile of tood
func (f *File) CreateFile(dir string) error {
	file, err := os.Create(dir + "/" + f.Name)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

// FindFromCurrentDirectory by filename
func (f *File) FindFromCurrentDirectory() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		path := dir + "/" + f.Name
		if err := f.SetFile(path); err == nil {
			return nil
		}

		if dir == "/" {
			return nil
		}

		dir = filepath.Dir(dir)
	}
}

// SetFile from path
func (f *File) SetFile(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}

	f.path = path
	return nil
}