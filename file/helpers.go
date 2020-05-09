package file

import (
	"os"
	"path/filepath"
)

const (
	// TodoFileName is name of todo collection file
	todoFileName string = ".todo"

	// TodoFilePermission set read permission
	todoFilePermission os.FileMode = 0644
)

// FindTodoFile from current directory
func FindTodoFile() *File {
	file := FindFromCurrentDirectory(todoFileName)
	if file != nil {
		file.Permission = todoFilePermission
	}

	return file
}

// FindFromDirectory by filename
func FindFromDirectory(name string, dir string) *File {
	for {
		path := dir + "/" + name
		if exist, _ := IsExist(path); exist {
			file := &File{Name: name, path: path}

			return file
		}

		if dir == "/" {
			return nil
		}

		dir = filepath.Dir(dir)
	}
}

// FindFromCurrentDirectory by filename
func FindFromCurrentDirectory(name string) *File {
	dir, err := os.Getwd()
	if err != nil {
		panic("Failed to get current path, please check permissions")
	}

	return FindFromDirectory(name, dir)
}

// IsExist check from path param
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	return os.IsNotExist(err) == false, err
}

// CreateTodoFile to dir param
func CreateTodoFile(dir string) error {
	return CreateFile(todoFileName, dir)
}

// CreateFile by name and dir
func CreateFile(name string, dir string) error {
	file, err := os.Create(dir + "/" + name)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
