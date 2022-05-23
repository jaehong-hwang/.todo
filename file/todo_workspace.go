package file

type TodoWorkspace struct {
	*Fileinfo
}

const (
	// TODO_FILE_NAME is name of todo main list file
	TODO_FILE_NAME string = "todo.json"

	HISTORY_DIRECTORY string = "histories"
)

func NewTodoWorkspace(fileinfo *Fileinfo) *TodoWorkspace {
	return &TodoWorkspace{
		fileinfo,
	}
}
