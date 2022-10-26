package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// LIST := BracketOpen [EXPRS] BracketClose
func acceptList(a auditor) (ast.List, bool) {
	if a.isNot(token.BracketOpen) {
		return ast.List{}, false
	}

	n := ast.List{
		Opener: a.Read(),
		Values: acceptExprsUntil(a, token.BracketClose),
		Closer: a.expect(token.BracketClose),
	}

	return n, true
}
