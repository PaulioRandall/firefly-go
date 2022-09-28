package compiler

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	//"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Compile(tr tokenreader.TokenReader) ([]ast.Node, error) {
	var (
		nodes []ast.Node
		n     ast.Node
		e     error
	)

	for tr.More() {
		tk := tr.Peek()

		switch {
		case tk.Type.IsLiteral():
			n, e = expectLiteral(tr)
		default:
			return nil, errors.New("Unexpected token") // TODO: Make proper error
		}

		if e != nil {
			return nil, e // TODO: Wrap error
		}

		nodes = append(nodes, n)
	}

	return nodes, nil
}

func acceptLiteral(tr tokenreader.TokenReader) (ast.Node, bool) {
	if !tr.More() {
		return nil, false
	}

	if tr.Peek().Type.IsLiteral() {
		return ast.MakeLiteral(tr.Read()), true
	}

	return nil, false
}

func expectLiteral(tr tokenreader.TokenReader) (ast.Node, error) {
	if !tr.More() {
		return nil, errors.New("Unexpected end of file") // TODO: Make proper error
	}

	tk := tr.Read()
	if !tk.Type.IsLiteral() {
		return nil, errors.New("Expected literal") // TODO: Make proper error
	}

	return ast.MakeLiteral(tk), nil
}
