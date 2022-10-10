package container

type Consumer[T any] interface {
	More() bool
	Empty() bool
	Len() int
	First() (T, bool)
	Take() (T, bool)
	Return(T)
}

type Provider[T any] interface {
	More() bool
	Empty() bool
	Len() int
	Last() (T, bool)
	Add(T)
	Reclaim() (T, bool)
}

type Queue[T any] interface {
	Consumer[T]
	Provider[T]
}

type LinkedQueue[T any] struct {
	front *node[T]
	back  *node[T]
	size  int
}

func (q LinkedQueue[T]) More() bool {
	return q.front != nil
}

func (q LinkedQueue[T]) Empty() bool {
	return q.front == nil
}

func (q LinkedQueue[T]) Len() int {
	return q.size
}

func (q LinkedQueue[T]) First() (T, bool) {
	if q.More() {
		return q.front.v, true
	}

	var v T
	return v, false
}

func (q LinkedQueue[T]) Last() (T, bool) {
	if q.More() {
		return q.back.v, true
	}

	var v T
	return v, false
}

func (q *LinkedQueue[T]) Add(v T) {
	n := &node[T]{
		v: v,
	}

	if q.Empty() {
		q.front = n
		q.back = n
		q.size = 1
		return
	}

	n.prev = q.back
	q.back.next = n
	q.back = n
	q.size++
}

func (q *LinkedQueue[T]) Return(v T) {
	n := &node[T]{
		v: v,
	}

	if q.Empty() {
		q.front = n
		q.back = n
		q.size = 1
		return
	}

	n.next = q.front
	q.front.prev = n
	q.front = n
	q.size++
}

func (q *LinkedQueue[T]) Take() (T, bool) {
	v, ok := q.First()
	if !ok {
		return v, false
	}

	q.front = q.front.next
	q.size--

	if q.front == nil {
		q.back = nil
	}

	return v, true
}

func (q *LinkedQueue[T]) Reclaim() (T, bool) {
	v, ok := q.Last()
	if !ok {
		return v, false
	}

	q.back = q.back.prev
	q.size--

	if q.back == nil {
		q.front = nil
	}

	return v, true
}
