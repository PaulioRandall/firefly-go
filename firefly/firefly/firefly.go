package firefly

import (
	"io/ioutil"

	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/cleaner"
	"github.com/PaulioRandall/firefly-go/firefly/grouper"
	"github.com/PaulioRandall/firefly-go/firefly/parser"
	"github.com/PaulioRandall/firefly-go/firefly/runner"
	"github.com/PaulioRandall/firefly-go/firefly/scanner"
	"github.com/PaulioRandall/firefly-go/firefly/token"
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

func ParseFile(file string) (ast.Block, error) {

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
