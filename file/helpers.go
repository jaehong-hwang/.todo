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
	TODO_FILE_PERMISSION os.FileMode = 0755
)

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic("Failed to get current path, please check permissions")
	}

	return dir
}

func GetHomeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		panic("Failed to get current user")
	}

	return usr.HomeDir
}

// FindTodoWorkspace from current directory
func FindTodoWorkspace(dir string, increase bool) *TodoWorkspace {
	fileinfo := FindFromDirectory(TODO_DIRECTORY_NAME, dir, increase)
	if fileinfo != nil {
		fileinfo.Permission = TODO_FILE_PERMISSION
	}

	workspace, err := NewTodoWorkspace(fileinfo)
	if err != nil {
		panic(err)
	}

	return workspace
}

// FindTodoSystemFile from home directory
func FindTodoSystemFile() *File {
	dir := GetHomeDirectory()
	fileinfo := FindFromDirectory(TODO_SYSTEM_FILE_NAME, dir, false)
	if fileinfo != nil {
		fileinfo.Permission = TODO_FILE_PERMISSION
	}

	return &File{
		fileinfo,
	}
}

// FindFromDirectory by filename
func FindFromDirectory(name string, dir string, increase bool) *Fileinfo {
	fromDir := dir
	for {
		path := dir + "/" + name
		if exist, _ := IsExist(path); exist {
			fileinfo := &Fileinfo{
				Name: name,
				path: path,
				directory: dir,
			}

			return fileinfo
		}

		if dir == "/" || increase == false {
			return &Fileinfo{
				Name:      name,
				path: fromDir + "/" + name,
				directory: fromDir,
			}
		}

		dir = filepath.Dir(dir)
	}
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

// CreateDirectory by path
func CreateDirectory(path string) error {
	return os.Mkdir(path, TODO_FILE_PERMISSION)
}
