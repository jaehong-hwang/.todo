package file

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	// TODO_DIRECTORY_NAME is name of todo collection directory
	TODO_DIRECTORY_NAME string = ".todo"

	// TODO_SYSTEM_FILE_NAME is name of configuration file
	TODO_SYSTEM_FILE_NAME string = ".todo.system"

	// TODO_FILE_PERMISSION set read permission
	TODO_FILE_PERMISSION os.FileMode = 0644
)

// FindTodoFile from current directory
func FindTodoFile(increase bool) *File {
	file := FindFromCurrentDirectory(TODO_DIRECTORY_NAME, increase)
	if file != nil {
		file.Permission = TODO_FILE_PERMISSION
	}

	return file
}

// FindTodoFile from current directory
func FindTodoFileWithDirectory(dir string, increase bool) *File {
	file := FindFromDirectory(TODO_DIRECTORY_NAME, dir, increase)
	if file != nil {
		file.Permission = TODO_FILE_PERMISSION
	}

	return file
}

// FindTodoSystemFile from home directory
func FindTodoSystemFile() *File {
	usr, err := user.Current()
	if err != nil {
		panic("Failed to get current user")
	}

	file := FindFromDirectory(TODO_SYSTEM_FILE_NAME, usr.HomeDir, false)
	if file != nil {
		file.Permission = TODO_FILE_PERMISSION
	}

	return file
}

// FindFromDirectory by filename
func FindFromDirectory(name string, dir string, increase bool) *File {
	fromDir := dir
	for {
		path := dir + "/" + name
		if exist, _ := IsExist(path); exist {
			file := &File{
				Name:      name,
				path:      path,
				directory: dir,
			}

			return file
		}

		if dir == "/" || increase == false {
			return &File{
				Name:      name,
				path:      fromDir + "/" + name,
				directory: fromDir,
			}
		}

		dir = filepath.Dir(dir)
	}
}

// FindFromCurrentDirectory by filename
func FindFromCurrentDirectory(name string, increase bool) *File {
	dir, err := os.Getwd()
	if err != nil {
		panic("Failed to get current path, please check permissions")
	}

	return FindFromDirectory(name, dir, increase)
}

// IsExist check from path param
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	return os.IsNotExist(err) == false, err
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

func CreateIfNotExists(name string, dir string) error {
	exists, err := IsExist(dir + "/" + name)
	if exists == false {
		return CreateFile(name, dir)
	} else if err != nil {
		return err
	}

	return nil
}
