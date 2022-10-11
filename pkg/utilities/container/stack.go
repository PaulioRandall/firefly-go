package container

import (
	"errors"
)

type Stack[T any] interface {
	More() bool
	Empty() bool
	Top() T
	Push(T)
	Pop() T
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

func (st *LinkedStack[T]) Top() T {
	if st.More() {
		return st.top.v
	}
	panic(errors.New("Stack is empty")) // TODO: Replace with bool
}

func (st *LinkedStack[T]) Push(v T) {
	st.top = &node[T]{
		v:    v,
		next: st.top,
	}
}

func (st *LinkedStack[T]) Pop() T {
	v := st.Top()
	st.top = st.top.next
	return v
}
