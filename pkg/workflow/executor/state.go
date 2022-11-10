package executor

import (
	"fmt"
)

type exeState struct {
	variables map[string]any
	spells    map[string]any
	e         error
}

func newState() *exeState {
	return &exeState{
		variables: map[string]any{},
	}
}

func (state *exeState) getVariable(name string) any {
	v, _ := state.variables[name]
	return v
}

func (state *exeState) getSpell(name string) any {
	v, _ := state.spells[name]
	return v
}

func (state *exeState) getError() error {
	return state.e
}

func (state *exeState) hasError() bool {
	return state.e != nil
}

func (state *exeState) setVariable(name string, value any) {
	state.variables[name] = value
}

func (state *exeState) setError(e error) {
	state.e = e
}

func (state *exeState) Println() {
	fmt.Println("** State **")
	for k, v := range state.variables {
		fmt.Println("\t"+k+": ", v)
	}
}
