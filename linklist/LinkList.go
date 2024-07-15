package linklist

import "generic-collections/interfaces"

type LinkList[T any] struct {
	interfaces.ICollection[T]
	Head  *Node[T]
	Tail  *Node[T]
	count int
}

var _ interfaces.ICollection[any] = (*LinkList[any])(nil)
