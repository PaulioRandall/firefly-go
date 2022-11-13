package memory

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/spells"
)

type Memory struct {
	Variables map[string]any
	Spells    map[string]spells.Spell
	Error     error
}

func (mem *Memory) Var(name string) any {
	return mem.Variables[name]
}

func (mem *Memory) Spell(name string) spells.Spell {
	return mem.Spells[name]
}

func NewMemory() *Memory {
	return &Memory{
		Variables: map[string]any{},
		Spells:    spells.All,
	}
}

func (mem *Memory) Println() {
	fmt.Println("** Memory **")
	for k, v := range mem.Variables {
		fmt.Println("\t"+k+": ", v)
	}
}
