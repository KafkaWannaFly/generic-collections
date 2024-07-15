package btree

import "generic-collections/interfaces"

type BTree[T any] struct {
	interfaces.ICollection[T]

	root  *Node[T]
	count int
}

var _ interfaces.ICollection[any] = (*BTree[any])(nil)
