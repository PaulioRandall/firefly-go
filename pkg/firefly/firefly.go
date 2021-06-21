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
	sr := token.NewStringScrollReader(runes)
	lxs, e := scanner.ScanAll(sr)
	if e != nil {
		return nil, e
	}

	lr := token.NewSliceLexemeReader(lxs)
	stmts, e := grouper.GroupAll(lr)
	if e != nil {
		return nil, e
	}

	pr := token.NewProgramReader(stmts)
	stmts, e = cleaner.CleanAll(pr)
	if e != nil {
		return nil, e
	}

	pr = token.NewProgramReader(stmts)
	program, e := parser.ParseAll(pr)
	if e != nil {
		return nil, e
	}

	return program, nil
}
