// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ASTWriter = inout.Writer[ast.Node]

func Parse(r TokenReader, w ASTWriter) (e error) {
	a := newAuditor(r)

	defer func() {
		if v := recover(); v != nil {
			e = err.Wrap(v.(error), "Recovered from parse fail")
		}
	}()

	return parseRootStatements(a, w)
}

func parseRootStatements(a *auditor, w ASTWriter) error {
	a.accept(token.Terminator)

	for a.more() {
		n := expectStatement(a)
		if e := w.Write(n); e != nil {
			return err.Wrap(e, "Parser failed to parse statements in the root scope")
		}
	}

	return nil
}
