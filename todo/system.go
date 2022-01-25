package todo

import (
	"encoding/json"

	"github.com/jaehong-hwang/todo/file"
)

type Directories []string
type System struct {
	file        *file.File
	Directories Directories `json:"directories"`
	Author      Author      `json:"author"`
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
	for _, dir := range s.Directories {
		if dir == directory {
			return nil
		}
	}

	s.Directories = append(s.Directories, directory)
	return s.Save()
}

func (s *System) GetDirectoryJson() (string, error) {
	b, err := json.Marshal(s.Directories)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *System) Save() error {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return err
	}

	s.file.FillContent(string(jsonData))

	return nil
}
