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
}

// IsExists todo file in current directory
func (f *File) IsExists() bool {
	return f.path != "" && f.path != "/"
}

// GetContent from todo file
func (f *File) GetContent() (string, error) {
	if f.IsExists() == false {
		return "", errors.New(f.Name + " file not found")
	}

	content, err := ioutil.ReadFile(f.path)
	return string(content), err
}

// FillContent to todo file
func (f *File) FillContent(content string) error {
	if f.IsExists() == false {
		return errors.New(f.Name + " file not found")
	}

	return ioutil.WriteFile(f.path, []byte(content), f.Permission)
}

// SetPath from path
func (f *File) SetPath(path string) error {
	if _, err := IsExist(path); err != nil {
		return err
	}

	f.path = path
	return nil
}
