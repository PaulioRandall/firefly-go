package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	UnexpectedEOF   = err.Trackable("Unexpected end of file")
	UnexpectedToken = err.Trackable("Unexpected token")

	MissingVar  = err.Trackable("Missing variable")
	MissingExpr = err.Trackable("Missing expression")
	MissingStmt = err.Trackable("Missing statement")
	MissingEnd  = err.Trackable("Missing end")
)
