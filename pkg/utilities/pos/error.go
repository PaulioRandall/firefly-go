package pos

import (
	"fmt"
)

type posErr struct {
	pos   Pos
	cause error
	msg   string
}

func ErrorAt(pos Pos, msg string, args ...interface{}) *posErr {
	return ErrorFor(pos, nil, msg, args...)
}

func ErrorFor(pos Pos, cause error, msg string, args ...interface{}) *posErr {
	return &posErr{
		pos:   pos,
		cause: cause,
		msg:   fmt.Sprintf(msg, args...),
	}
}

func (e posErr) Error() string {
	return e.msg
}

func (e posErr) Unwrap() error {
	return e.cause
}

func (e posErr) Pos() Pos {
	return e.pos
}
