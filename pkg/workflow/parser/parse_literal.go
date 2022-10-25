package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var MissingLiteral = err.Trackable("Missing literal")

// LITERAL := True | False | Number | String

func acceptLiteral(a auditor) ast.Expr {

	switch a.Peek() {
	case token.True:
	case token.False:
	case token.Number:
	case token.String:
	default:
		return nil
	}

	return ast.Literal{
		Token: a.Read(),
	}
}

func expectLiteral(a auditor) ast.Expr {
	if n := acceptLiteral(a); n != nil {
		return n
	}
	panic(MissingLiteral.Trackf("Expected literal but got %s", a.Peek().String()))
}
