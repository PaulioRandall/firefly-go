package token

import (
	"errors"
)

// EOF is returned by readers if the end of a stream (file) was reached. It may
// be compared to read errors to determine if EOF was the cause.
var EOF = errors.New("EOF")
