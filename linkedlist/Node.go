package linkedlist

import (
	"github.com/KafkaWannaFly/generic-collections/interfaces"
	"github.com/KafkaWannaFly/generic-collections/utils"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

var _ interfaces.IHashCoder = (*Node[any])(nil)

func NewNode[T any]() *Node[T] {
	return &Node[T]{}
}

func NodeOf[T any](value T) *Node[T] {
	return &Node[T]{Value: value}
}

func (receiver *Node[T]) HashCode() string {
	return utils.HashCodeOf(receiver.Value)
}

// Clone creates a new Node with the same value as the receiver Node. However, didn't copy the Next field.
// Return a new Node
func (receiver *Node[T]) Clone() *Node[T] {
	node := NewNode[T]()
	node.Value = receiver.Value

	return node
}
