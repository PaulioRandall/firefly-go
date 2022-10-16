package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	UnexpectedEOF   = err.New("Unexpected end of file")
	UnexpectedToken = err.New("Unexpected token")

	MissingVar  = err.New("Missing variable")
	MissingExpr = err.New("Missing expression")
	MissingStmt = err.New("Missing statement")
	MissingEnd  = err.New("Missing end")
)
