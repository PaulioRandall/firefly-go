package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/container"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

type panicPosReader[T pos.Positioned] struct {
	reader Reader[T]
	buffer container.Queue[T]
	prev   T
}

func NewPanicPosReader[T pos.Positioned](r Reader[T]) *panicPosReader[T] {
	return &panicPosReader[T]{
		reader: r,
		buffer: &container.LinkedQueue[T]{},
	}
}

func (a panicPosReader[T]) Prev() T {
	return a.prev
}

func (a *panicPosReader[T]) More() bool {
	return a.buffer.More() || a.reader.More()
}

func (a *panicPosReader[T]) Peek() T {
	a.loadBuffer()

	if tk, ok := a.buffer.First(); ok {
		return tk
	}

	_, to := a.prev.Where()
	panic(err.WrapPos(EOF, to, "Failed to peek from buffer"))
}

func (a *panicPosReader[T]) Read() T {
	a.loadBuffer()

	if tk, ok := a.buffer.Take(); ok {
		a.prev = tk
		return tk
	}

	_, to := a.prev.Where()
	panic(err.WrapPos(EOF, to, "Failed to read from buffer"))
}

func (a *panicPosReader[T]) Putback(v T) {
	a.buffer.Return(v)
}

func (a *panicPosReader[T]) loadBuffer() {
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
