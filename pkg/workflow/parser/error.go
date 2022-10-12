package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
)

var (
	MissingVar  = err.New("Missing variable")
	MissingExpr = err.New("Missing expression")
	MissingStmt = err.New("Missing statement")
	MissingEnd  = err.New("Missing end")
)
