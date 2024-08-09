package btree

import "github.com/KafkaWannaFly/generic-collections/interfaces"

type Node[T any] struct {
	interfaces.IComparer[Node[T]]
	interfaces.IHashCoder

	value T
	left  *Node[T]
	right *Node[T]
}

var _ interfaces.IComparer[Node[any]] = (*Node[any])(nil)
var _ interfaces.IHashCoder = (*Node[any])(nil)
