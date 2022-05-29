package file

import "os"

type Fileinfo struct {
	Name string
	Permission os.FileMode
	path string
	directory string
}

// IsExist todo file in current directory
func (f *Fileinfo) IsExist() bool {
	isExistsFlag, _ := IsExist(f.path)
	return isExistsFlag
}

// CreateIfNotExist will make file if is not exists
func (f *Fileinfo) CreateIfNotExist() error {
	return CreateIfNotExists(f.Name, f.directory)
}

// GetDirectory from todo file
func (f *Fileinfo) GetDirectory() string {
	return f.directory
}

// Remove current file
func (f *Fileinfo) Remove() error {
	return os.Remove(f.path)
}

// Find from current directory
func (f *Fileinfo) Find(name string) *Fileinfo {
	return FindFromDirectory(name, f.path, false)
}