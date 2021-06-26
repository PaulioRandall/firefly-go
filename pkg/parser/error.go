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

func panicParseErr(cause error, msg string, args ...interface{}) {
	panic(newParseErr(cause, msg, args...))
}

func newParseErr(cause error, msg string, args ...interface{}) *parseErr {
	return &parseErr{
		msg:   fmt.Sprintf(msg, args...),
		cause: cause,
	}
}
