package workflow

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/workflow/cleaner"
	"github.com/PaulioRandall/firefly-go/pkg/workflow/executor"
	parser "github.com/PaulioRandall/firefly-go/pkg/workflow/parser2"
	"github.com/PaulioRandall/firefly-go/pkg/workflow/scanner"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

type ReaderOfRunes = inout.Reader[rune]
type WriterOfNodes = inout.Writer[ast.Node]

func Run(r ReaderOfRunes) (int, error) {

	var (
		tks        []token.Token
		nodes      []ast.Node
		e          error
		parseError = func(e error) error {
			return err.Wrap(e, "Workflow failed to parse scroll")
		}
	)

	if tks, e = scan(r); e != nil {
		return 1, parseError(e)
	} else if tks == nil {
		return 0, nil
	}

	if tks, e = clean(tks); e != nil {
		return 1, parseError(e)
	} else if tks == nil {
		return 0, nil
	}

	if nodes, e = parse(tks); e != nil {
		return 1, parseError(e)
	} else if nodes == nil {
		return 0, nil
	}

	state := executor.NewState()
	executor.Execute(state, nodes)
	return state.GetExitCode(), state.GetError()
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

func parse(tks []token.Token) ([]ast.Node, error) {
	r := inout.NewListReader(tks)
	w := inout.NewListWriter[ast.Node]()

	if e := parser.Parse(r, w); e != nil {
		return nil, err.Wrap(e, "Failed to parse AST")
	}

	if w.Empty() {
		return nil, nil
	}

	return w.List(), nil
}
