package errors

var errors = map[string]string{
    "todo_already_exists": "todo collection already exists",
    "todo_doesnt_exists": "todo dosen't exists, you should run todo init",
    "unexpected_state": "${state} is unexpected state. todo have 3 state ex. wait, work, done",
}
