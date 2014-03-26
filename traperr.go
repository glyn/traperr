package traperr

import (
	"runtime"
)

type Err interface {

	// Returns the error message.
	Error() string

	// Returns the stack trace of the error.
	StackTrace() string
}

// Returns an Err containing the given message.
func New(message string) Err {
	var stack [4096]byte

	runtime.Stack(stack[:], false)

	pc, file, line, _ := runtime.Caller(1)

	return &errImpl{message, stack[:], pc, file, line}
}

type errImpl struct {
	message string
	stackTrace []byte
	pc uintptr
	file string
	line int
}

func (e *errImpl) Error() string {
	return e.message
}

func (e *errImpl) StackTrace() string {
	return string(e.stackTrace)
}
