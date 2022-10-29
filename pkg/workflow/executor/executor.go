package executor

import (
	"log"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
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

	for _, n := range nodes {
		executeNode(state, n)

		if state.hasError() {
			break
		}
	}
}
