package stack

import (
	"errors"
)

type Stack[T any] interface {
	Empty() bool
	Top() T
	Push(T)
	Pop() T
}

type LinkedStack[T any] struct {
	top *node[T]
}

type node[T any] struct {
	v    T
	next *node[T]
}

func (st LinkedStack[T]) Empty() bool {
	return st.top == nil
}

func (st *LinkedStack[T]) Top() T {
	if st.Empty() {
		panic(errors.New("Stack is empty"))
	}
	return st.top.v
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
