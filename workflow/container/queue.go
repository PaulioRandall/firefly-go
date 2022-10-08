package container

import (
	"errors"
)

type Consumer[T any] interface {
	More() bool
	Empty() bool
	Next() T
	Take() T
	Push(T)
}

type Producer[T any] interface {
	More() bool
	Empty() bool
	Last() T
	Add(T)
	TakeBack() T
}

type Queue[T any] interface {
	Consumer[T]
	Producer[T]
}

type LinkedQueue[T any] struct {
	front *node[T]
	back  *node[T]
}

func (q LinkedQueue[T]) More() bool {
	return q.front != nil
}

func (q LinkedQueue[T]) Empty() bool {
	return q.front == nil
}

func (q LinkedQueue[T]) Next() T {
	if q.More() {
		return q.front.v
	}
	panic(errors.New("Queue is empty"))
}

func (q LinkedQueue[T]) Last() T {
	if q.More() {
		return q.back.v
	}
	panic(errors.New("Queue is empty"))
}

func (q *LinkedQueue[T]) Add(v T) {
	n := &node[T]{
		v: v,
	}

	if q.Empty() {
		q.front = n
		q.back = n
		return
	}

	n.prev = q.back
	q.back.next = n
	q.back = n
}

func (q *LinkedQueue[T]) Push(v T) {
	n := &node[T]{
		v: v,
	}

	if q.Empty() {
		q.front = n
		q.back = n
		return
	}

	n.next = q.front
	q.front.prev = n
	q.front = n
}

func (q *LinkedQueue[T]) Take() T {
	v := q.Next()
	q.front = q.front.next

	if q.front == nil {
		q.back = nil
	}

	return v
}

func (q *LinkedQueue[T]) TakeBack() T {
	v := q.Last()
	q.back = q.back.prev

	if q.back == nil {
		q.front = nil
	}

	return v
}
