package set

import "generic-collections/interfaces"

type Set[T any] struct {
	interfaces.ICollection[T]

	elements map[string]T
	count    int
}
