package inout

type listReader[T any] struct {
	idx  int
	data []T
	prev T
}

func NewListReader[T any](list []T) *listReader[T] {
	return &listReader[T]{
		data: list,
	}
}

func (lr listReader[T]) More() bool {
	return lr.idx < len(lr.data)
}

func (lr *listReader[T]) Peek() (T, error) {
	if lr.More() {
		return lr.data[lr.idx], nil
	}
	var zero T
	return zero, EOF
}

func (lr *listReader[T]) Read() (T, error) {
	if lr.More() {
		v := lr.data[lr.idx]
		lr.idx++
		lr.prev = v
		return v, nil
	}

	var zero T
	return zero, EOF
}

func (lr listReader[T]) Prev() T {
	return lr.prev
}
