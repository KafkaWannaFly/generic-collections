package queue

import "generic-collections/interfaces"

type Queue[T any] struct {
	interfaces.ICollection[T]
	elements []T
	count    int
}

var _ interfaces.ICollection[any] = (*Queue[any])(nil)
