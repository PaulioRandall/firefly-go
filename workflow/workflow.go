package workflow

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/steps/aligner"
	"github.com/PaulioRandall/firefly-go/workflow/steps/compiler"
	"github.com/PaulioRandall/firefly-go/workflow/steps/formaliser"
	"github.com/PaulioRandall/firefly-go/workflow/steps/rinser"
	"github.com/PaulioRandall/firefly-go/workflow/steps/scanner"
	"github.com/PaulioRandall/firefly-go/workflow/token"
	"github.com/PaulioRandall/firefly-go/workflow/tokenreader"
)

type RuneReader interface {
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}

/*
type NodeOutput interface {
	Write(ast.Node) error
}
*/

func Parse(r RuneReader) ([]ast.Node, error) {

	var (
		tks    []token.Token
		e      error
		failed = func(e error) error {
			return fmt.Errorf("Failed to parse scroll: %w", e)
		}
	)

	if tks, e = scan(r); e != nil {
		return nil, failed(e)
	} else if tks == nil {
		return nil, nil
	}

	if tks, e = rinse(tks); e != nil {
		return nil, failed(e)
	} else if tks == nil {
		return nil, nil
	}

	if tks, e = align(tks); e != nil {
		return nil, failed(e)
	} else if tks == nil {
		return nil, nil
	}

	// TODO: Refactor next
	// TODO: Think about combining aligner & formaliser
	tr := tokenreader.FromList(tks...)
	tks = formaliser.Formalise(tr)

	tr = tokenreader.FromList(tks...)
	nodes, e := compiler.Compile(tr)
	if e != nil {
		return nil, fmt.Errorf("Failed to scan scroll: %w", e)
	}

	return nodes, nil
}

func scan(r RuneReader) ([]token.Token, error) {
	w := inout.NewListWriter[token.Token]()

	if e := scanner.Scan(r, w); e != nil {
		return nil, fmt.Errorf("Failed to scan scroll: %w", e)
	}

	if w.Empty() {
		return nil, nil
	}
	return w.List(), nil
}

func rinse(tks []token.Token) ([]token.Token, error) {
	r := inout.NewListReader(tks)
	w := inout.NewListWriter[token.Token]()

	if e := rinser.Rinse(r, w); e != nil {
		return nil, fmt.Errorf("Failed to rinse scroll: %w", e)
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
		return nil, fmt.Errorf("Failed to align scroll: %w", e)
	}
	return w.List(), nil
}
