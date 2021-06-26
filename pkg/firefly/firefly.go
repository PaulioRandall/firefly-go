package firefly

import (
	"io/ioutil"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/cleaner"
	"github.com/PaulioRandall/firefly-go/pkg/grouper"
	"github.com/PaulioRandall/firefly-go/pkg/parser"
	"github.com/PaulioRandall/firefly-go/pkg/runner"
	"github.com/PaulioRandall/firefly-go/pkg/scanner"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func RunFile(file string) error {

	p, e := ParseFile(file)
	if e != nil {
		return e
	}

	in := runner.NewInterpreter(p)
	in.Exe()
	return in.ExeErr()
}

func ParseFile(file string) (ast.Program, error) {

	data, e := ioutil.ReadFile(file)
	if e != nil {
		return nil, e
	}

	runes := []rune(string(data))
	runeReader := token.NewRuneReader(runes)
	lxs, e := scanner.ScanAll(runeReader)
	if e != nil {
		return nil, e
	}

	lexemeReader := token.NewLexReader(lxs)
	stmts, e := grouper.GroupAll(lexemeReader)
	if e != nil {
		return nil, e
	}

	stmtReader := token.NewStmtReader(stmts)
	stmts, e = cleaner.CleanAll(stmtReader)
	if e != nil {
		return nil, e
	}

	stmtReader = token.NewStmtReader(stmts)
	program, e := parser.ParseAll(stmtReader)
	if e != nil {
		return nil, e
	}

	return program, nil
}
