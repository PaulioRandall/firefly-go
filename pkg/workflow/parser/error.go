package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/err"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

var (
	UnexpectedEOF   = auditor.UnexpectedEOF
	UnexpectedToken = auditor.UnexpectedToken

	MissingVar  = err.New("Missing variable")
	MissingExpr = err.New("Missing expression")
	MissingStmt = err.New("Missing statement")
	MissingEnd  = err.New("Missing end")
)
