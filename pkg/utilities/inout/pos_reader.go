package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/container"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

type posReader[T pos.Positioned] struct {
	reader Reader[T]
	buffer container.Queue[T]
	prev   T
}

func NewPosReader[T pos.Positioned](r Reader[T]) *posReader[T] {
	return &posReader[T]{
		reader: r,
		buffer: &container.LinkedQueue[T]{},
	}
}

func (a posReader[T]) Prev() T {
	return a.prev
}

func (a *posReader[T]) More() bool {
	return a.buffer.More() || a.reader.More()
}

func (a *posReader[T]) Peek() T {
	a.loadBuffer()

	if tk, ok := a.buffer.First(); ok {
		return tk
	}

	_, to := a.prev.Where()
	panic(err.WrapPos(EOF, to, "Failed to peek from buffer"))
}

func (a *posReader[T]) Read() T {
	a.loadBuffer()

	if tk, ok := a.buffer.Take(); ok {
		a.prev = tk
		return tk
	}

	_, to := a.prev.Where()
	panic(err.WrapPos(EOF, to, "Failed to read from buffer"))
}

func (a *posReader[T]) Putback(v T) {
	a.buffer.Return(v)
}

func (a *posReader[T]) loadBuffer() {
	if a.buffer.More() {
		return
	}

	v, e := a.reader.Read()
	if e != nil {
		_, to := a.prev.Where()
		panic(err.WrapPos(e, to, "Failed to read from reader"))
	}

	a.buffer.Add(v)
}
