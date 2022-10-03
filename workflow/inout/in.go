package inout

import (
	"errors"
)

var EOF = errors.New("End of file")

type inputList[T any] struct {
	idx  int
	data []T
}

func FromList[T any](list []T) inputList[T] {
	return inputList[T]{
		data: list,
	}
}

func (in inputList[T]) More() bool {
	return in.idx < len(in.data)
}

func (in *inputList[T]) Peek() (T, error) {
	if in.More() {
		return in.data[in.idx], nil
	}
	var zero T
	return zero, EOF
}

func (in *inputList[T]) Read() (T, error) {
	if in.More() {
		v := in.data[in.idx]
		in.idx++
		return v, nil
	}

	var zero T
	return zero, EOF
}
