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

func NewPos(from pos.Pos, m string) *fireflyError {
	e := New(m)
	e.from = from
	return e
}

func NewPosf(from pos.Pos, m string, args ...any) *fireflyError {
	e := Newf(m, args...)
	e.from = from
	return e
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

func WrapPos(cause error, from pos.Pos, m string) *fireflyError {
	e := New(m)
	e.cause = cause
	e.from = from
	return e
}

func WrapPosf(cause error, from pos.Pos, m string, args ...any) *fireflyError {
	e := Newf(m, args...)
	e.cause = cause
	e.from = from
	return e
}

func WrapRange(cause error, from, to pos.Pos, m string) *fireflyError {
	e := New(m)
	e.cause = cause
	e.from = from
	e.to = to
	return e
}

func WrapRangef(cause error, from, to pos.Pos, m string, args ...any) *fireflyError {
	e := Newf(m, args...)
	e.cause = cause
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
