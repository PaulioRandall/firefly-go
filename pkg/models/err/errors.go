package err

import (
	"io"
)

var (
	EOF             = Wrap(io.EOF, io.EOF.Error())
	UnexpectedEOF   = New("Unexpected end of file")
	UnexpectedToken = New("Unexpected token")
)
