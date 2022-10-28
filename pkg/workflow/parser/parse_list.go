package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrMissingBracketClose = err.Trackable("Missing closing bracket")

	ErrBadList = err.Trackable("Failed to parse list")
)

// LIST := BracketOpen [EXPRS] BracketClose
func acceptList(a auditor) (ast.List, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadList.Wrap(e, "Bad list syntax")
	})

	if a.isNot(token.BracketOpen) {
		return ast.List{}, false
	}

	n := ast.List{
		Opener: a.Read(),
		Values: acceptExprsUntil(a, token.BracketClose),
		Closer: parseBracketClose(a),
	}

	return n, true
}

func parseBracketClose(a auditor) token.Token {
	defer wrapPanic(func(e error) error {
		return ErrMissingBracketClose.Wrap(e, "Expected closing bracket")
	})

	return a.expect(token.BracketClose)
}
