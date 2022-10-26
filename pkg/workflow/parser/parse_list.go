package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectList(a auditor) ast.List {
	return ast.List{
		Opener: a.expect(token.BracketOpen),
		Values: acceptExprsUntil(a, token.BracketClose),
		Closer: a.expect(token.BracketClose),
	}
}
