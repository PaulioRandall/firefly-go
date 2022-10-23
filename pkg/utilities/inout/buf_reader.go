package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/container"
)

type bufReader[T any] struct {
	reader Reader[T]
	buffer container.Queue[T]
	prev   T
}

func NewBufReader[T any](r Reader[T]) *bufReader[T] {
	return &bufReader[T]{
		reader: r,
		buffer: &container.LinkedQueue[T]{},
	}
}

func (r bufReader[T]) Prev() T {
	return r.prev
}

func (r *bufReader[T]) More() bool {
	return r.buffer.More() || r.reader.More()
}

func (r *bufReader[T]) Peek() (T, error) {
	var zero T

	if e := r.buff(); e != nil {
		return zero, ErrRead.Track(e, "Peeking failed")
	}

	r.prev, _ = r.buffer.First()
	return r.prev, nil
}

func (r *bufReader[T]) Read() (T, error) {
	var zero T

	if e := r.buff(); e != nil {
		return zero, ErrRead.Track(e, "Reading failed")
	}

	r.prev, _ = r.buffer.Take()
	return r.prev, nil
}

func (r *bufReader[T]) Putback(v T) {
	r.buffer.Return(v)
}

func (r *bufReader[T]) buff() error {
	if r.buffer.More() {
		return nil
	}

	v, e := r.reader.Read()
	if e != nil {
		return ErrReadDelegate.Track(e, "Buffering failed")
	}

	r.buffer.Add(v)
	return nil
}
