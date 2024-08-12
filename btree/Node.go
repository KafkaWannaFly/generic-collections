package btree

import "github.com/KafkaWannaFly/generic-collections/interfaces"

type Node[T any] struct {
	interfaces.IHashCoder

	value T
	left  *Node[T]
	right *Node[T]
}

var _ interfaces.IHashCoder = (*Node[any])(nil)
