package file

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	// TodoFileName is name of todo collection file
	todoFileName string = ".todo"

	// TodoFileName is name of todo collection file
	todoSystemFileName string = ".todo.system"

	// TodoFilePermission set read permission
	todoFilePermission os.FileMode = 0644
)

// FindTodoFile from current directory
func FindTodoFile() *File {
	file := FindFromCurrentDirectory(todoFileName, false)
	if file != nil {
		file.Permission = todoFilePermission
	}

	return file
}

// FindTodoFile from current directory
func FindTodoFileWithDirectory(dir string) *File {
	file := FindFromDirectory(todoFileName, dir, false)
	if file != nil {
		file.Permission = todoFilePermission
	}

	return file
}

// FindTodoSystemFile from home directory
func FindTodoSystemFile() *File {
	usr, err := user.Current()
	if err != nil {
		panic("Failed to get current user")
	}

	file := FindFromDirectory(todoSystemFileName, usr.HomeDir, false)
	if file != nil {
		file.Permission = todoFilePermission
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
