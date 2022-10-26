package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// MAP := BraceOpen {MAP_ENTRY} BraceClose
func acceptMap(a auditor) (ast.Map, bool) {
	if a.isNot(token.BraceOpen) {
		return ast.Map{}, false
	}

	n := ast.Map{
		Opener:  a.Read(),
		Entries: parseMapEntries(a),
		Closer:  a.expect(token.BraceClose),
	}

	return n, true
}

// MAP_ENTRY := EXPR Colon EXPR
func parseMapEntries(a auditor) []ast.MapEntry {
	var entries []ast.MapEntry

	for a.isNot(token.BraceClose) {
		entry := parseMapEntry(a)
		entries = append(entries, entry)

		if !a.accept(token.Comma) {
			break
		}
	}

	return entries
}

// MAP_ENTRY := EXPR Colon EXPR
func parseMapEntry(a auditor) ast.MapEntry {
	key := expectExpression(a)
	a.expect(token.Colon)
	value := expectExpression(a)

	return ast.MapEntry{
		Key:   key,
		Value: value,
	}
}
