package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectMap(a auditor) ast.Map {
	return ast.Map{
		Opener:  a.expect(token.BraceOpen),
		Entries: expectMapEntries(a),
		Closer:  a.expect(token.BraceClose),
	}
}

func expectMapEntries(a auditor) []ast.MapEntry {
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
