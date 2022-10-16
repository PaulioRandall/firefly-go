// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ReaderOfTokens = inout.Reader[token.Token]
type WriterOfNodes = inout.Writer[ast.Node]

type BufReaderOfTokens = inout.PosReader[token.Token]

func Parse(r ReaderOfTokens, w WriterOfNodes) (e error) {
	br := BufReaderOfTokens(inout.NewPosReader[token.Token](r))

	defer func() {
		if v := recover(); v != nil {
			e = err.Wrap(v.(error), "Recovered from parse fail")
		}
	}()

	return parseRootStatements(br, w)
}

func parseRootStatements(r BufReaderOfTokens, w WriterOfNodes) error {
	accept(r, token.Terminator)

	for r.More() {
		n := expectStatement(r)
		if e := w.Write(n); e != nil {
			return err.Wrap(e, "Parser failed to parse statements in the root scope")
		}
	}

	return nil
}
