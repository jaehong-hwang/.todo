package todo

import (
	"encoding/json"

	"github.com/jaehong-hwang/todo/file"
)

type Directory struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Directories []Directory
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

func (s *System) AddDirectory(path string) error {
	for _, dir := range s.Directories {
		if dir.Path == path {
			return nil
		}
	}

	s.Directories = append(s.Directories, Directory{
		Name: "",
		Path: path,
	})
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
