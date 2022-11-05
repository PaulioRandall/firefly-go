package executor

import (
	"log"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
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

	state := NewState()
	exeNodes(state, nodes)
	return state.getExitCode(), state.getError()
}
