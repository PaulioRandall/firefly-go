package executor

type exeState struct {
	variables map[string]any
	e         error
}

func NewState() *exeState {
	return &exeState{
		variables: map[string]any{},
	}
}

func (state *exeState) setError(e error) {
	state.e = e
}

func (state *exeState) hasError() bool {
	return state.e != nil
}

func (state *exeState) getVariable(name string) (any, bool) {
	if v, ok := state.variables[name]; ok {
		return v, true
	}
	return nil, false
}

func (state *exeState) setVariable(name string, value any) {
	state.variables[name] = value
}
