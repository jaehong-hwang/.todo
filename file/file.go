package file

import (
	"errors"
	"io/ioutil"
	"os"
)

// File management struct
type File struct {
	Name       string
	Permission os.FileMode
	path       string
	directory  string
}

// IsExist todo file in current directory
func (f *File) IsExist() bool {
	isExistsFlag, _ := IsExist(f.path)
	return isExistsFlag
}

// CreateIfNotExist will make file if is not exists
func (f *File) CreateIfNotExist() error {
	if f.IsExist() == true {
		return nil
	} else {
		return CreateFile(f.Name, f.directory)
	}
}

// GetContent from todo file
func (f *File) GetContent() (string, error) {
	if f.IsExist() == false {
		return "", errors.New(f.Name + " file not found")
	}

	content, err := ioutil.ReadFile(f.path)
	return string(content), err
}

// FillContent to todo file
func (f *File) FillContent(content string) error {
	if f.IsExist() == false {
		return errors.New(f.Name + " file not found")
	}

	return ioutil.WriteFile(f.path, []byte(content), f.Permission)
}
