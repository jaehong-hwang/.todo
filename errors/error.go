package errors

import "strings"

type TodoError struct {
	Code    string
	Message string
}

func (e *TodoError) Error() string {
	return e.Message
}

func New(code string) error {
	return &TodoError{
		Code:    code,
		Message: errors[code],
	}
}

func NewWithParam(code string, param map[string]string) error {
	errorMessage := errors[code]
	for key, val := range param {
		errorMessage = strings.Replace(errorMessage, "${"+key+"}", val, -1)
	}
	return &TodoError{
		Code:    code,
		Message: errorMessage,
	}
}
