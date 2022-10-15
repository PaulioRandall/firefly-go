package asttest

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func Variables(tks ...token.Token) []ast.Variable {
	var nodes []ast.Variable

	for _, tk := range tks {
		nodes = append(nodes, Variable(tk))
	}

	return nodes
}

func Expressions(tks ...token.Token) []ast.Expr {
	var nodes []ast.Expr

	for _, tk := range tks {
		nodes = append(nodes, Literal(tk))
	}

	return nodes
}
