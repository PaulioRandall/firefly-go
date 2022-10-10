package workflow

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/utilities/inout"

	"github.com/PaulioRandall/firefly-go/models/ast"
	"github.com/PaulioRandall/firefly-go/models/token"

	"github.com/PaulioRandall/firefly-go/workflow/steps/aligner"
	"github.com/PaulioRandall/firefly-go/workflow/steps/cleaner"
	"github.com/PaulioRandall/firefly-go/workflow/steps/parser"
	"github.com/PaulioRandall/firefly-go/workflow/steps/scanner"
	"github.com/PaulioRandall/firefly-go/workflow/steps/terminator"
)

type RuneReader = inout.Reader[rune]
type NodeWriter = inout.Writer[ast.Node]

func Parse(r RuneReader, w NodeWriter) error {

	var (
		tks    []token.Token
		e      error
		failed = func(e error) error {
			return fmt.Errorf("Failed to parse scroll: %w", e)
		}
	)

	if tks, e = scan(r); e != nil {
		return failed(e)
	} else if tks == nil {
		return nil
	}

	if tks, e = clean(tks); e != nil {
		return failed(e)
	} else if tks == nil {
		return nil
	}

	if tks, e = terminate(tks); e != nil {
		return failed(e)
	}

	if tks, e = align(tks); e != nil {
		return failed(e)
	}

	return parse(tks, w)
}

func scan(r RuneReader) ([]token.Token, error) {
	w := inout.NewListWriter[token.Token]()

	if e := scanner.Scan(r, w); e != nil {
		return nil, fmt.Errorf("Failed to scan tokens: %w", e)
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
		return nil, fmt.Errorf("Failed to remove redundant tokens: %w", e)
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
		return nil, fmt.Errorf("Failed to inline comma separated values: %w", e)
	}
	return w.List(), nil
}

func terminate(tks []token.Token) ([]token.Token, error) {
	r := inout.NewListReader(tks)
	w := inout.NewListWriter[token.Token]()

	if e := terminator.Terminate(r, w); e != nil {
		return nil, fmt.Errorf("Failed to convert newlines to terminators: %w", e)
	}
	return w.List(), nil
}

func parse(tks []token.Token, w NodeWriter) error {
	r := inout.NewListReader(tks)

	if e := parser.Parse(r, w); e != nil {
		return fmt.Errorf("Failed to parse AST: %w", e)
	}

	return nil
}
