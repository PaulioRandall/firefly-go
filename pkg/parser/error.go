package parser

import (
	"fmt"
)

type ParseError interface {
	error
	Cause() error
}

type parseErr struct {
	msg   string
	cause error
}

func (e parseErr) Error() string {
	return e.msg
}

func (e parseErr) Cause() error {
	return e.cause
}

func newThenPanic(msg string, args ...interface{}) {
	panic(newErr(msg, args...))
}

func wrapThenPanic(cause error, msg string, args ...interface{}) {
	panic(wrapErr(cause, msg, args...))
}

func newErr(msg string, args ...interface{}) parseErr {
	return parseErr{
		msg: fmt.Sprintf(msg, args...),
	}
}

func wrapErr(cause error, msg string, args ...interface{}) parseErr {
	return parseErr{
		msg:   fmt.Sprintf(msg, args...),
		cause: cause,
	}
}
