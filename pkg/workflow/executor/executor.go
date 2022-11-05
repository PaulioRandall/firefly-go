package executor

import (
	"log"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func Execute(state *exeState, nodes []ast.Node) {

	defer func() {
		v := recover()
		if v == nil {
			return
		}

		e, ok := v.(error)
		if !ok {
			log.Fatalf("Recovered from panic that was not an error: %v", v)
		}

		state.setError(e)
	}()

	exeNodes(state, nodes)
}
