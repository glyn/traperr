package traperr

type Err interface {
	Error() string
}

// Returns an err containing the given message.
func New(message string) Err {
	return &errImpl{message}
}

type errImpl struct {
	message string
}

// Returns the message associated with the err.
func (e *errImpl) Error() string {
	return e.message
}
