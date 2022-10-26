package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func acceptList(a auditor) (ast.List, bool) {
	var zero ast.List

	if !a.is(token.BracketOpen) {
		return zero, false
	}

	n := ast.List{
		Opener: a.Read(),
		Values: acceptExprsUntil(a, token.BracketClose),
		Closer: a.expect(token.BracketClose),
	}

	return n, true
}
