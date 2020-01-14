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

// IsExist todo file in current directory
func (f *File) IsExist() bool {
	isExistsFlag, _ := IsExist(f.path)
	return isExistsFlag
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
