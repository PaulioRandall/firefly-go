// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ReaderOfTokens = inout.Reader[token.Token]
type WriterOfNodes = inout.Writer[ast.Node]

type tokenAuditor = *auditor.Auditor[token.Token]

func Parse(r ReaderOfTokens, w WriterOfNodes) (e error) {
	a := tokenAuditor(auditor.NewAuditor[token.Token](r))

	defer func() {
		if v := recover(); v != nil {
			e = err.Wrap(v.(error), "Recovered from parse fail")
		}
	}()

	return parseRootStatements(a, w)
}

func parseRootStatements(a tokenAuditor, w WriterOfNodes) error {
	accept(a, token.Terminator)

	for a.More() {
		n := expectStatement(a)
		if e := w.Write(n); e != nil {
			return err.Wrap(e, "Parser failed to parse statements in the root scope")
		}
	}

	return nil
}
