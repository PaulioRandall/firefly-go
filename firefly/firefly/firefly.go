// Package firefly provides convenience functions for parsing and interpreting
// firefly scrolls.
//
// If more control is needed the underlying packages can be used directly and
// this package can serve as an example of usage.
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

// RunFile parses and then executes the specified file returning an error if
// one is encountered.
func RunFile(file string) error {

	p, e := ParseFile(file)
	if e != nil {
		return e
	}

	in := runner.NewInterpreter(p)
	in.Exe()
	return in.ExeErr()
}

// ParseFile parses the specified file returning an AST block or an error if
// one is encountered. It performs the sfollowing steps in order:
//		1. scanning: scans tokens from the file
//		2. grouping: groups tokens into statements
//		3. cleaning: cleans these statements, e.g. removes whitespace etc
//		4. parsing:  parses the statements into Abstract Syntax Trees
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
