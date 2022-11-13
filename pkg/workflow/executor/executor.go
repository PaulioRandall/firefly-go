package executor

import (
	"log"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

func Execute(nodes []ast.Node) (exitCode int, e error) {

	defer func() {
		v := recover()
		if v == nil {
			return
		}

		var ok bool
		if e, ok = v.(error); !ok {
			log.Fatalf("Recovered from panic that was not an error: %v", v)
		}
	}()

	mem := memory.NewMemory()
	exeNodes(mem, nodes)

	mem.Println()

	if mem.Error != nil {
		return 1, mem.Error
	}

	return 0, nil
}
