// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ASTWriter = inout.Writer[ast.Node]

func Parse(r TokenReader, w ASTWriter) (e error) {
	a := newAuditor(r)

	defer func() {
		if v := recover(); v != nil {
			e = v.(error) // TODO: Wrap error
		}
	}()

	for r.More() {
		n := parseNext(a)
		e := w.Write(n)
		if e != nil {
			panic(e)
		}
	}

	return nil
}

func parseNext(a *auditor) (n ast.Node) {
	switch {
	case a.accept(token.Var):
		n = parseStartingWithVariable(a, a.prev)

	case a.isNext(token.If):
		n = parseIf(a)

	default:
		panic(UnexpectedToken)
	}

	if n == nil {
		panic(errors.New("Sanity check! Nil Node should never appear"))
	}

	a.expect(token.Terminator)
	return n
}

func parseStartingWithVariable(a *auditor, first token.Token) ast.Node {
	if a.isNext(token.Comma) || a.isNext(token.Assign) {
		a.putback(first)
		return parseAssignment(a)
	}

	panic(UnexpectedToken)
}
