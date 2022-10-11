package err

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

type FireflyError interface {
	error
	Unwrap() error
	Where() (pos.Pos, pos.Pos)
}

type fireflyError struct {
	msg   string
	cause error
	from  pos.Pos
	to    pos.Pos
}

func New(m string) *fireflyError {
	return &fireflyError{
		msg: m,
	}
}

func Newf(m string, args ...any) *fireflyError {
	return &fireflyError{
		msg: fmt.Sprintf(m, args...),
	}
}

func Wrap(cause error, m string) *fireflyError {
	e := New(m)
	e.cause = cause
	return e
}

func Wrapf(cause error, m string, args ...any) *fireflyError {
	e := Newf(m, args...)
	e.cause = cause
	return e
}

func (e *fireflyError) SetWhere(from, to pos.Pos) *fireflyError {
	e.from = from
	e.to = to
	return e
}

func (e *fireflyError) Error() string {
	return e.msg
}

func (e *fireflyError) Unwrap() error {
	return e.cause
}

func (e *fireflyError) Where() (pos.Pos, pos.Pos) {
	return e.from, e.to
}
