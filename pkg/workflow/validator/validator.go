// Package validator validates abstract syntax trees are valid statements
package validator

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	//"github.com/PaulioRandall/firefly-go/pkg/models/token"

	//"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ReaderOfNodes = inout.Reader[ast.Node]
type WriterOfNodes = inout.Writer[ast.Node]

func Validate(r ReaderOfNodes, w WriterOfNodes) (e error) {
	// TODO
	for r.More() {
		n, _ := r.Read()
		w.Write(n)
	}
	return nil
}
