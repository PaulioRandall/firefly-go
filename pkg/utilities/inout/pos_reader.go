package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
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
		panic(ErrReadDelegate.WrapPos(e, to, "Peeking failed"))
	}

	return v
}

func (r *posReader[T]) Read() T {
	v, e := r.BufReader.Read()

	if e != nil {
		_, to := r.Prev().Where()
		panic(ErrReadDelegate.WrapPos(e, to, "Reading failed"))
	}

	return v
}
