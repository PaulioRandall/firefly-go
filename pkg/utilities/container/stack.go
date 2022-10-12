package container

type Stack[T any] interface {
	More() bool
	Empty() bool
	Top() (T, bool)
	Push(T)
	Pop() (T, bool)
}

type LinkedStack[T any] struct {
	top *node[T]
}

func (st LinkedStack[T]) More() bool {
	return st.top != nil
}

func (st LinkedStack[T]) Empty() bool {
	return st.top == nil
}

func (st *LinkedStack[T]) Top() (T, bool) {
	if st.More() {
		return st.top.v, true
	}

	var zero T
	return zero, false
}

func (st *LinkedStack[T]) Push(v T) {
	st.top = &node[T]{
		v:    v,
		next: st.top,
	}
}

func (st *LinkedStack[T]) Pop() (T, bool) {
	if v, ok := st.Top(); ok {
		st.top = st.top.next
		return v, true
	}

	var zero T
	return zero, false
}
