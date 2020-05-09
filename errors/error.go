package errors

type TodoError struct {
	Code string
	Message string
}

func (e *TodoError) Error() string {
	return e.Message
}

func New(code string) error {
	return &TodoError{
		Code: code,
		Message: errors[code],
	}
}
