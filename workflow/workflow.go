package workflow

import (
	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/readers/runereader"
	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/steps/aligner"
	"github.com/PaulioRandall/firefly-go/workflow/steps/compiler"
	"github.com/PaulioRandall/firefly-go/workflow/steps/formaliser"
	"github.com/PaulioRandall/firefly-go/workflow/steps/rinser"
	"github.com/PaulioRandall/firefly-go/workflow/steps/scanner"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Parse(rr runereader.RuneReader) ([]ast.Node, error) {

	tks, e := scanner.ScanAll(rr)
	if e == err.EOF {
		return nil, nil
	} else if e != nil {
		return nil, err.AtPos(token.Pos{}, e, "Failed to scan scroll")
	}

	tr := tokenreader.FromList(tks...)
	tks, e = rinser.RinseAll(tr)
	if e != nil {
		return nil, e // TODO: wrap error
	}

	tr = tokenreader.FromList(tks...)
	tks = aligner.AlignAll(tr)

	tr = tokenreader.FromList(tks...)
	tks = formaliser.Formalise(tr)

	tr = tokenreader.FromList(tks...)
	nodes, e := compiler.Compile(tr)
	if e != nil {
		return nil, e // TODO: wrap error
	}

	return nodes, nil
}
