package err

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

// trackableError is an error that is tracked through it's error message.
//
// This is so that specific error type can be tracked without the underlying
// cause interfering with equality.
type trackableError struct {
	msg   string
	cause error
}

// Trackable creates a new template error from which specific errors can be
// constructed.
//
// Call the receiving Track function to create specific instantiations with a
// cause. The template error must be package scoped such that tests and
// consuming packages can access the template and check for equality.
func Trackable(m string) *trackableError {
	return &trackableError{
		msg: m,
	}
}

// Track returns a trackable error with a cause.
//
// errors.Is will return true if the receiving and resultant error are the
// inputs thus allowing a fairly specific error to be tracked while keeping
// the cause irrelevant.
func (e trackableError) Track(
	cause error,
	msg string,
	args ...any,
) *trackableError {
	e.cause = Wrapf(cause, msg, args...)
	return &e
}

func (e trackableError) TrackPos(
	cause error,
	from pos.Pos,
	msg string,
	args ...any,
) *trackableError {
	e.cause = WrapPosf(cause, from, msg, args...)
	return &e
}

func (e *trackableError) Error() string {
	return e.msg
}

func (e *trackableError) Unwrap() error {
	return e.cause
}

// Is compares only the error message for trackable errors.
func (e *trackableError) Is(target error) bool {
	if v, ok := target.(*trackableError); ok {
		return v.msg == e.msg
	}
	return false
}
