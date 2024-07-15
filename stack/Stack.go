package stack

import "generic-collections/interfaces"

type Stack[T any] struct {
	interfaces.ICollection[T]
	elements []T
	count    int
}

var _ interfaces.ICollection[any] = (*Stack[any])(nil)
