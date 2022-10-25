package err

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

// trackableError is an error that can be identified via errors.Is function.
//
// This is so errors of varying generality can be tracked without needing to
// create large numbers of specific error types or have the underlying cause or
// other error member values interfere with equality.
type trackableError struct {
	id    int
	msg   string
	cause error
}

var idPool = 0

// Trackable creates a new template error from which specific errors can be
// constructed.
//
// Call the receiving Track or Wrap functions to create specific
// instantiations. The template error must be package scoped such that tests
// and consuming packages can access the template and check for equality.
//
// The error message should be general enough to cover all use cases but no
// more. Create trackable error templates with varying message precision if
// multiple levels of tracking are required. There is no extension mechanism
// but trackable errors are errors so multiple trackable errors may be nested.
// For your own sanity, take care to always wrap specific errors within general
// ones, not the other way around.
func Trackable(m string) *trackableError {
	idPool++ // Not thread safe, does it need to be?

	return &trackableError{
		id:  idPool,
		msg: m,
	}
}

// Track returns a new trackable error with the same private ID as the
// receiving error.
//
// A non-trackable error will be created with your message then wrapped by the
// general trackable error so that all messages are preserved and a nice
// readable stack trace can by constructed easily.
//
// errors.Is will return true if the both errors passed to it are trackable
// errors with the same ID. The only way to create two trackable errors with
// the same ID is to use the Track or Wrap functions. Copying the error value
// is pointless as the initial general message is static.
func (e trackableError) Track(
	msg string,
) *trackableError {
	e.cause = New(msg)
	return &e
}

// Trackf is the same as Track with extra option for error meessage formatting.
func (e trackableError) Trackf(
	msg string,
	args ...any,
) *trackableError {
	e.cause = Newf(msg, args...)
	return &e
}

// TrackPos is the same as Track with extra option for a file position.
func (e trackableError) TrackPos(
	from pos.Pos,
	msg string,
) *trackableError {
	e.cause = NewPos(from, msg)
	return &e
}

// TrackPosf is the same as TrackPos with extra option for error meessage
// formatting.
func (e trackableError) TrackPosf(
	from pos.Pos,
	msg string,
	args ...any,
) *trackableError {
	e.cause = NewPosf(from, msg, args...)
	return &e
}

// Wrap is the same as Track but wraps a pre-constructed underlying cause.
//
// The cause will be wrapped once as a non-trackable error type with your
// message then wrapped by the general trackable error so that all
// messages are preserved and a nice readable stack trace can by constructed
// easily.
func (e trackableError) Wrap(
	cause error,
	msg string,
) *trackableError {
	e.cause = Wrap(cause, msg)
	return &e
}

// Trackf is the same as Wrap with extra option for error meessage formatting.
func (e trackableError) Wrapf(
	cause error,
	msg string,
	args ...any,
) *trackableError {
	e.cause = Wrapf(cause, msg, args...)
	return &e
}

// TrackPos is the same as Wrap with extra option for a file position.
func (e trackableError) WrapPos(
	cause error,
	from pos.Pos,
	msg string,
) *trackableError {
	e.cause = WrapPos(cause, from, msg)
	return &e
}

// TrackPosf is the same as WrapPos with extra option for error meessage
// formatting.
func (e trackableError) WrapPosf(
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

// Unwrap will always return a non-nil value for errors created using the
// Track or Wrap functions.
func (e *trackableError) Unwrap() error {
	return e.cause
}

// Is compares using a private hidden ID which is assigned only when Trackable
// is called and copied derived errors when the receiving Track or Wrap
// functions are called.
func (e *trackableError) Is(target error) bool {
	if v, ok := target.(*trackableError); ok {
		return v.id == e.id
	}
	return false
}
