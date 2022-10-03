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

type RuneInput interface {
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}

/*
type NodeOutput interface {
	Write(ast.Node) error
}
*/

func Parse(in RuneInput) ([]ast.Node, error) {

	var (
		tks    []token.Token
		e      error
		failed = func(e error) error {
			return fmt.Errorf("Failed to parse scroll: %w", e)
		}
	)

	if tks, e = scan(in); e != nil {
		return nil, failed(e)
	} else if tks == nil {
		return nil, nil
	}

	if tks, e = rinse(tks); e != nil {
		return nil, failed(e)
	} else if tks == nil {
		return nil, nil
	}

	// TODO: Refactor next
	// TODO: Think about combining aligner & formaliser
	tr := tokenreader.FromList(tks...)
	tks = aligner.AlignAll(tr)

	tr = tokenreader.FromList(tks...)
	tks = formaliser.Formalise(tr)

	tr = tokenreader.FromList(tks...)
	nodes, e := compiler.Compile(tr)
	if e != nil {
		return nil, fmt.Errorf("Failed to scan scroll: %w", e)
	}

	return nodes, nil
}

func scan(in RuneInput) ([]token.Token, error) {
	out := inout.ToList[token.Token]()

	if e := scanner.Scan(in, &out); e != nil {
		return nil, fmt.Errorf("Failed to scan scroll: %w", e)
	}

	if out.Empty() {
		return nil, nil
	}
	return out.List(), nil
}

func rinse(tks []token.Token) ([]token.Token, error) {
	in := inout.FromList(tks)
	out := inout.ToList[token.Token]()

	if e := rinser.Rinse(&in, &out); e != nil {
		return nil, fmt.Errorf("Failed to rinse scroll: %w", e)
	}

	if out.Empty() {
		return nil, nil
	}
	return out.List(), nil
}
