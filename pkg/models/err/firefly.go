package err

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/pos"
)

type FireflyError struct {
	Message string
	Cause   error
	From    pos.Pos
	To      pos.Pos
}

func New(m string) *FireflyError {
	return &FireflyError{
		Message: m,
	}
}

func Newf(m string, args ...any) *FireflyError {
	return &FireflyError{
		Message: fmt.Sprintf(m, args...),
	}
}

func Wrap(cause error, m string) *FireflyError {
	e := New(m)
	e.Cause = cause
	return e
}

func Wrapf(cause error, m string, args ...any) *FireflyError {
	e := Newf(m, args...)
	e.Cause = cause
	return e
}

func (e *FireflyError) SetWhere(from, to pos.Pos) *FireflyError {
	e.From = from
	e.To = to
	return e
}

func (e *FireflyError) Error() string {
	return e.Message
}

func (e *FireflyError) Unwrap() error {
	return e.Cause
}

func (e *FireflyError) Where() (pos.Pos, pos.Pos) {
	return e.From, e.To
}
