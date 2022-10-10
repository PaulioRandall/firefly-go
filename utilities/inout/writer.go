package inout

type listWriter[T any] struct {
	data []T
}

func NewListWriter[T any]() *listWriter[T] {
	return &listWriter[T]{}
}

func (w *listWriter[T]) Write(v T) error {
	w.data = append(w.data, v)
	return nil
}

func (w *listWriter[T]) WriteMany(v ...T) error {
	w.data = append(w.data, v...)
	return nil
}

func (w listWriter[T]) List() []T {
	return w.data
}

func (w listWriter[T]) Empty() bool {
	return len(w.data) == 0
}
