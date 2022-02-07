package errors

var errors = map[string]string{
	"todo_already_exists":  "todo collection already exists",
	"todo_doesnt_exists":   "todo dosen't exists, you should run todo init",
	"unexpected_state":     "${state} is unexpected state. todo have 3 state ex. wait, work, done",
	"message_required":     "message is required field",
	"todo_empty":           "todo is empty. you can add to do thing with `todo add` command",
	"todo_id_not_found":    "ID ${id} is not exists.",
	"label_already_exists": "${label} is already exists label",
	"label_not_found":      "${label} label not found",
	"sort_method_invalid":  "${sort} is invalid order method, you can sort only `regist-date`, `due-date`, `level`. and default value `regist-date`",
}
