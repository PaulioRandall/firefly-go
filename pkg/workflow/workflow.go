package workflow

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/workflow/aligner"
	"github.com/PaulioRandall/firefly-go/pkg/workflow/cleaner"
	"github.com/PaulioRandall/firefly-go/pkg/workflow/parser"
	"github.com/PaulioRandall/firefly-go/pkg/workflow/scanner"
	"github.com/PaulioRandall/firefly-go/pkg/workflow/terminator"
)

type ReaderOfRunes = inout.Reader[rune]
type WriterOfNodes = inout.Writer[ast.Node]

func Parse(r ReaderOfRunes, w WriterOfNodes) error {

	var (
		tks        []token.Token
		e          error
		parseError = func(e error) error {
			return err.Wrap(e, "Workflow failed to parse scroll")
		}
	)

	if tks, e = scan(r); e != nil {
		return parseError(e)
	} else if tks == nil {
		return nil
	}

	if tks, e = clean(tks); e != nil {
		return parseError(e)
	} else if tks == nil {
		return nil
	}

	if tks, e = terminate(tks); e != nil {
		return parseError(e)
	}

	if tks, e = align(tks); e != nil {
		return parseError(e)
	}

	return parse(tks, w)
}

func scan(r ReaderOfRunes) ([]token.Token, error) {
	w := inout.NewListWriter[token.Token]()

	if e := scanner.Scan(r, w); e != nil {
		return nil, err.Wrap(e, "Workflow failed to scan tokens")
	}

	if w.Empty() {
		return nil, nil
	}

	return w.List(), nil
}

func clean(tks []token.Token) ([]token.Token, error) {
	r := inout.NewListReader(tks)
	w := inout.NewListWriter[token.Token]()

	if e := cleaner.Clean(r, w); e != nil {
		return nil, err.Wrap(e, "Workflow failed to clean tokens")
	}

	if w.Empty() {
		return nil, nil
	}

	return w.List(), nil
}

func align(tks []token.Token) ([]token.Token, error) {
	r := inout.NewListReader(tks)
	w := inout.NewListWriter[token.Token]()

	if e := aligner.Align(r, w); e != nil {
		return nil, err.Wrap(e, "Failed to align tokens")
	}

	return w.List(), nil
}

func terminate(tks []token.Token) ([]token.Token, error) {
	r := inout.NewListReader(tks)
	w := inout.NewListWriter[token.Token]()

	if e := terminator.Terminate(r, w); e != nil {
		return nil, err.Wrap(e, "Failed to convert newlines to terminators")
	}

	return w.List(), nil
}

func parse(tks []token.Token, w WriterOfNodes) error {
	r := inout.NewListReader(tks)

	if e := parser.Parse(r, w); e != nil {
		return err.Wrap(e, "Failed to parse AST")
	}

	return nil
}
