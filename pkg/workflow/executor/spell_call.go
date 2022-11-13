package executor

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

var (
	ErrUnknownSpell = err.Trackable("Unknown spell")
)

func invokeSpell(mem *memory.Memory, sp ast.SpellCall) []any {

	spell := mem.Spells[sp.Name]

	if spell == nil {
		panic(ErrUnknownSpell.Trackf("No spell found with name %q", sp.Name))
	}

	params := make([]any, len(sp.Params))
	for i, p := range sp.Params {
		params[i] = exeExpr(mem, p)
	}

	return spell(mem, params)
}
