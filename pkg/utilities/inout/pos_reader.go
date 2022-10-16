package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

type posReader[T pos.Positioned] struct {
	BufReader[T]
}

func NewPosReader[T pos.Positioned](r BufReader[T]) *posReader[T] {
	return &posReader[T]{
		BufReader: r,
	}
}

func (r *posReader[T]) Peek() T {
	v, e := r.BufReader.Peek()

	if e != nil {
		_, to := r.Prev().Where()
		panic(err.WrapPos(e, to, "Failed to peek"))
	}

	return v
}

func (r *posReader[T]) Read() T {
	v, e := r.BufReader.Read()

	if e != nil {
		_, to := r.Prev().Where()
		panic(err.WrapPos(e, to, "Failed to read"))
	}

	return v
}
