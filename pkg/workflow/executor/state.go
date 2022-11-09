package executor

import (
	"fmt"
)

type exeState struct {
	variables map[string]any
	e         error
}

func newState() *exeState {
	return &exeState{
		variables: map[string]any{},
	}
}

func (state *exeState) setError(e error) {
	state.e = e
}

func (state *exeState) getError() error {
	return state.e
}

func (state *exeState) hasError() bool {
	return state.e != nil
}

func (state *exeState) getVariable(name string) any {
	if v, ok := state.variables[name]; ok {
		return v
	}
	return nil
}

func (state *exeState) setVariable(name string, value any) {
	state.variables[name] = value
}

func (state *exeState) Println() {
	fmt.Println("** State **")
	for k, v := range state.variables {
		fmt.Println("\t"+k+": ", v)
	}
}
