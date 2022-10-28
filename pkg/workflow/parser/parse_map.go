package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadMap      = err.Trackable("Failed to parse map")
	ErrBadMapEntry = err.Trackable("Failed to parse map entry")
)

// MAP := BraceOpen MAP_ENTRIES BraceClose
func acceptMap(a auditor) (ast.Map, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadMap.Wrap(e, "Bad map syntax")
	})

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

// MAP_ENTRIES := [MAP_ENTRY {Comma MAP_ENTRY}]
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
	defer wrapPanic(func(e error) error {
		return ErrBadMapEntry.Wrap(e, "Bad map entry")
	})

	key := expectExpression(a)
	a.expect(token.Colon)
	value := expectExpression(a)

	return ast.MapEntry{
		Key:   key,
		Value: value,
	}
}
