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

type PosReaderOfTokens = inout.PosReader[token.Token]

var ErrParsing = err.Trackable("Parsing failed")

func Parse(r ReaderOfTokens, w WriterOfNodes) (e error) {
	br := inout.NewBufReader[token.Token](r)
	pr := inout.NewPosReader[token.Token](br)
	a := auditor{
		r: pr,
	}

	defer func() {
		if v := recover(); v != nil {
			e = ErrParsing.Track(v.(error), "Recovered from parse fail")
		}
	}()

	return parseRootStatements(a, w)
}

func parseRootStatements(a auditor, w WriterOfNodes) error {
	a.accept(token.Terminator)

	for a.More() {
		n := expectStatement(a)
		if e := w.Write(n); e != nil {
			return ErrParsing.Track(e, "Parser failed to parse statements in the root scope")
		}
	}

	return nil
}
