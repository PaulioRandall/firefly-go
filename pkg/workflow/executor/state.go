package executor

type exeState struct {
	e error
}

func (state *exeState) setError(e error) {
	state.e = e
}

func (state *exeState) hasError() bool {
	return state.e != nil
}
