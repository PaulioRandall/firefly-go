package container

type node[T any] struct {
	v    T
	next *node[T]
	prev *node[T]
}
