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

	tokenOut := inout.ToList[token.Token]()

	if e := scanner.Scan(in, &tokenOut); e != nil {
		return nil, fmt.Errorf("Failed to scan scroll: %w", e)
	}

	tks := tokenOut.List()
	if len(tks) == 0 {
		return nil, nil
	}

	tr := tokenreader.FromList(tks...)
	tks = rinser.RinseAll(tr)

	tr = tokenreader.FromList(tks...)
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
