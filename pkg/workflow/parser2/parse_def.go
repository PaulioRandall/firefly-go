package parser2

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

// TODO:
// - spell
// - func
// - proc

var (
	ErrBadDef = err.Trackable("Failed to define procedure or function")
)

func acceptDef(a auditor) (ast.Stmt, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadDef.Wrap(e, "Bad procedure or function")
	})

	if !a.accept(token.Def) {
		return nil, false
	}

	return nil, false
}
