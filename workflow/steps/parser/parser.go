// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type NodeWriter = inout.Writer[ast.Node]

func Parse(r TokenReader, w NodeWriter) (e error) {
	a := &auditor{
		TokenReader: r,
	}

	defer func() {
		if v := recover(); v != nil {
			e = v.(error) // TODO: Wrap error
		}
	}()

	for r.More() {
		var n ast.Node

		switch {
		case a.acceptIf(token.IsLiteral):
			n = ast.Literal{Token: a.get()}
		default:
			panic(errors.New("Unexpected token")) // TODO: Wrap error
		}

		if n == nil {
			panic(errors.New("Sanity check! Nil Node should never appear"))
		}

		if e = w.Write(n); e != nil {
			panic(e)
		}

		a.expect(token.Terminator)
	}

	return nil
}
