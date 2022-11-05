package asttest

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func ExprFor(tk token.Token) ast.Expr {
	if isLiteral(tk.TokenType) {
		return Literal(tk)
	} else if tk.TokenType == token.Identifier {
		return Variable(tk)
	}

	panic("asttest: unmanagedd token type")
}

func isLiteral(tt token.TokenType) bool {
	switch tt {
	case token.Number, token.String, token.True, token.False:
		return true
	default:
		return false
	}
}

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
		nodes = append(nodes, ExprFor(tk))
	}

	return nodes
}
