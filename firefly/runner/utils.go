package runner

import (
	"errors"
	"fmt"
)

func newBug(msg string, args ...interface{}) error {
	msg = "[BUG] " + msg
	return newError(msg, args...)
}

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
