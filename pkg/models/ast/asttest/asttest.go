package asttest

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func Vars(tks ...token.Token) []ast.Variable {
	var nodes []ast.Variable

	for _, tk := range tks {
		nodes = append(nodes, ast.MakeVariable(tk))
	}

	return nodes
}

func LitExprs(tks ...token.Token) []ast.Expr {
	var nodes []ast.Expr

	for _, tk := range tks {
		nodes = append(nodes, ast.MakeLiteral(tk))
	}

	return nodes
}
