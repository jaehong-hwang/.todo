package file

type TodoWorkspace struct {
	*Fileinfo
	mainfile *File
}

const (
	// TODO_FILE_NAME is name of todo main list file
	TODO_FILE_NAME string = "todo.json"

	HISTORY_DIRECTORY string = "histories"
)

func NewTodoWorkspace(fileinfo *Fileinfo) (*TodoWorkspace, error) {
	mainfile := &File{
		fileinfo.Find(TODO_FILE_NAME),
	}

	if fileinfo.IsExist() {
		err := mainfile.CreateIfNotExist()
		if err != nil {
			return nil, err
		}
	}

	return &TodoWorkspace{
		Fileinfo: fileinfo,
		mainfile: mainfile,
	}, nil
}

// GetContent from todo file
func (t *TodoWorkspace) GetContent() (string, error) {
	return t.mainfile.GetContent()
}

// FillContent to todo file
func (t *TodoWorkspace) FillContent(content string) error {
	return t.mainfile.FillContent(content)
}

// CreateIfNotExist for only todo_workspace
func (t *TodoWorkspace) CreateIfNotExist() error {
	if t.IsExist() == true {
		return nil
	}
	
	return CreateDirectory(t.path)
}
