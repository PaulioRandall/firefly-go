// Package validator validates abstract syntax trees are valid statements
package validator

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	//"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ReaderOfNodes = inout.Reader[ast.Node]
type WriterOfNodes = inout.Writer[ast.Node]

func Parse(r ReaderOfNodes, w WriterOfNodes) (e error) {
	return err.New("Not yet implemented")
}
