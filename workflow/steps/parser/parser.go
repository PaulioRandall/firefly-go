// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type ASTWriter = inout.Writer[ast.Node]

func Parse(r TokenReader, w ASTWriter) (e error) {
	a := &auditor{
		TokenReader: r,
	}

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
	case a.acceptIf(token.IsLiteral):
		n = ast.Literal{Token: a.get()}
	default:
		panic(errors.New("Unexpected token")) // TODO: Wrap error
	}

	if n == nil {
		panic(errors.New("Sanity check! Nil Node should never appear"))
	}

	a.expect(token.Terminator)
	return n
}
