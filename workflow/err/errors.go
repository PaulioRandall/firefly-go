package err

import (
	"errors"
)

var (
	EOF             = errors.New("End of file")
	UnexpectedEOF   = errors.New("Unexpected end of file")
	UnexpectedToken = errors.New("Unexpected token")
)
