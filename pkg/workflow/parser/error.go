package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	UnexpectedEOF   = auditor.UnexpectedEOF
	UnexpectedToken = auditor.UnexpectedToken

	MissingVar  = err.New("Missing variable")
	MissingExpr = err.New("Missing expression")
	MissingStmt = err.New("Missing statement")
	MissingEnd  = err.New("Missing end")
)
