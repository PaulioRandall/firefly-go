// Package validator validates abstract syntax trees are valid statements
package validator

import (
	//"errors"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	//"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ASTReader = inout.Reader[ast.Node]
type ASTWriter = inout.Writer[ast.Node]

func Parse(r ASTReader, w ASTWriter) (e error) {
	return nil
}
