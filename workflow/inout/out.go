package inout

type outputList[T any] struct {
	data []T
}

func ToList[T any]() outputList[T] {
	return outputList[T]{}
}

func (in *outputList[T]) Write(v ...T) error {
	in.data = append(in.data, v...)
	return nil
}

func (in outputList[T]) List() []T {
	return in.data
}

func (in outputList[T]) Empty() bool {
	return len(in.data) == 0
}
