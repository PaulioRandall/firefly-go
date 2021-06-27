package parser

import (
	"fmt"
)

// ParseError interface is implemented by all explicit parsing errors. Errors
// not satisfying this interface and likely software errors or bugs.
type ParseError interface {
	error
	Cause() error
}

type parseErr struct {
	msg   string
	cause error
}

// Error satisfies the error interface.
func (e parseErr) Error() string {
	if e.cause != nil {
		return e.cause.Error() + "\nWrapped by: " + e.msg
	}
	return e.msg
}

// Cause returns the underlying cause of this error.
func (e parseErr) Cause() error {
	return e.cause
}

func parsingPanic(cause error, msg string, args ...interface{}) {
	panic(newParseErr(cause, msg, args...))
}

func newParseErr(cause error, msg string, args ...interface{}) *parseErr {
	return &parseErr{
		msg:   fmt.Sprintf(msg, args...),
		cause: cause,
	}
}
