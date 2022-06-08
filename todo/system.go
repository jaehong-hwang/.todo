package todo

import (
	"encoding/json"

	"github.com/jaehong-hwang/todo/file"
)

type Directory struct {
	ID   string `json:"id"`
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

func (s *System) AddDirectory(newDir Directory) error {
	for _, dir := range s.Directories {
		if dir.Path == newDir.Path {
			return nil
		}
	}

	s.Directories = append(s.Directories, newDir)
	return s.Save()
}

func (s *System) RemoveDirectory(removePath string) error {
	for i, dir := range s.Directories {
		if dir.Path == removePath {
			s.Directories[i] = s.Directories[len(s.Directories)-1]
			s.Directories = s.Directories[:len(s.Directories)-1]
			break
		}
	}

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
