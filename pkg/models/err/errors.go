package err

import (
	"errors"
	"io"
)

var (
	EOF             = io.EOF
	UnexpectedEOF   = errors.New("Unexpected end of file")
	UnexpectedToken = errors.New("Unexpected token")
)
