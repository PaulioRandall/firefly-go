// Package parser parses a series of tokens into series of abstract syntax trees
package parser2

/*
import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ReaderOfTokens = inout.Reader[token.Token]
type WriterOfNodes = inout.Writer[ast.Node]

var (
	ErrParsing = err.Trackable("Failed to parse scroll")
	ErrWriting = err.Trackable("Failed to write output")
)

func Parse(r ReaderOfTokens, w WriterOfNodes) (e error) {
	br := inout.NewBufReader[token.Token](r)
	pr := inout.NewPosReader[token.Token](br)
	a := auditor{
		r: pr,
	}

	defer func() {
		if v := recover(); v != nil {
			e = ErrParsing.Wrap(v.(error), "Recovered from parse fail")
		}
	}()

	return parseRootStatements(a, w)
}

func parseRootStatements(a auditor, w WriterOfNodes) error {
	defer wrapPanic(func(e error) error {
		return ErrParsing.Wrap(e, "Failed to parse root statements")
	})

	a.accept(token.Terminator)

	for a.More() {
		n := expectStatement(a)
		if e := w.Write(n); e != nil {
			panic(ErrWriting.Wrap(e, "Couldn't write AST node to output"))
		}
	}

	return nil
}
*/
