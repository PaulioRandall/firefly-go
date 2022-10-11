package err

var (
	UnexpectedEOF   = New("Unexpected end of file")
	UnexpectedToken = New("Unexpected token")
)
