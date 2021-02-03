package todo

import (
	"encoding/json"

	"github.com/jaehong-hwang/todo/file"
)

type System struct {
	file        *file.File
	Directories []string `json:"directories"`
	Author      Author   `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewSystem(systemFile *file.File) *System {
	systemFile.CreateIfNotExist()

	system := System{
		file: systemFile,
	}

	if systemFile != nil {
		input, err := systemFile.GetContent()
		if err == nil {
			json.Unmarshal([]byte(input), &system)
		}
	}

	return &system
}

func (s *System) AddDirectory(directory string) error {
	s.Directories = append(s.Directories, directory)
	return s.Save()
}

func (s *System) Save() error {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return err
	}

	s.file.FillContent(string(jsonData))

	return nil
}
