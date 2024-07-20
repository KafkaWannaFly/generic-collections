package linkedlist

import (
	"generic-collections/interfaces"
	"generic-collections/utils"
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
