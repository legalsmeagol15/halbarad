package dependency

import "fmt"

type depError struct {
	source  Dependent
	message string
}

func newError(src Dependent, msg string, params ...any) depError {
	msg = fmt.Sprintf(msg, params)
	return depError{
		source:  src,
		message: msg,
	}
}
