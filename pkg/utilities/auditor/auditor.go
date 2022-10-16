package auditor

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/container"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type Auditor[T pos.Positioned] struct {
	reader inout.Reader[T]
	buffer container.Queue[T]
	prev   T
}

func NewAuditor[T pos.Positioned](r inout.Reader[T]) *Auditor[T] {
	return &Auditor[T]{
		reader: r,
		buffer: &container.LinkedQueue[T]{},
	}
}

func (a Auditor[T]) Prev() T {
	return a.prev
}

func (a *Auditor[T]) More() bool {
	return a.buffer.More() || a.reader.More()
}

func (a *Auditor[T]) Peek() T {
	a.loadBuffer()

	if tk, ok := a.buffer.First(); ok {
		return tk
	}

	_, to := a.prev.Where()
	panic(err.WrapPos(inout.EOF, to, "Failed to peek from buffer"))
}

func (a *Auditor[T]) Read() T {
	a.loadBuffer()

	if tk, ok := a.buffer.Take(); ok {
		a.prev = tk
		return tk
	}

	_, to := a.prev.Where()
	panic(err.WrapPos(inout.EOF, to, "Failed to read from buffer"))
}

func (a *Auditor[T]) Putback(v T) {
	a.buffer.Return(v)
}

func (a *Auditor[T]) loadBuffer() {
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
