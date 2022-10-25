package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// VAR := Identifier
func expectIdentifier(a auditor) ast.Expr {
	return ast.Variable{
		Identifier: a.expect(token.Identifier),
	}
}
