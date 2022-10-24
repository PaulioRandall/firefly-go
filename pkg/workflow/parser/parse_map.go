package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func parseMap(a auditor) ast.Map {
	return ast.Map{
		Opener:  a.expect(token.BraceOpen),
		Entries: parseMapEntries(a),
		Closer:  a.expect(token.BraceClose),
	}
}

func parseMapEntries(a auditor) []ast.MapEntry {
	var entries []ast.MapEntry

	for a.isNot(token.BraceClose) {
		key := expectExpression(a)
		a.expect(token.Colon)
		value := expectExpression(a)

		entry := ast.MapEntry{
			Key:   key,
			Value: value,
		}

		entries = append(entries, entry)

		if !a.accept(token.Comma) {
			break
		}
	}

	return entries
}
