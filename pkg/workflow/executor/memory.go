package executor

import (
	"fmt"
)

type Memory struct {
	Variables map[string]any
	Spells    map[string]any
	Error     error
}

func NewMemory() *Memory {
	return &Memory{
		Variables: map[string]any{},
		Spells:    map[string]any{},
	}
}

func (mem *Memory) Println() {
	fmt.Println("** Memory **")
	for k, v := range mem.Variables {
		fmt.Println("\t"+k+": ", v)
	}
}
