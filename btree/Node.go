package btree

import (
	"github.com/KafkaWannaFly/generic-collections/interfaces"
	"github.com/KafkaWannaFly/generic-collections/utils"
	"math"
)

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

var _ interfaces.IHashCoder = (*Node[interfaces.IHashCoder])(nil)

func NewNode[T any](value T, left *Node[T], right *Node[T]) *Node[T] {
	return &Node[T]{value: value, left: left, right: right}
}

func NewLeafNode[T any](value T) *Node[T] {
	return NewNode(value, nil, nil)
}

// Clone creates a new Node with the same value as the receiver Node.
// However, didn't copy the left and right fields.
func (receiver *Node[T]) Clone() *Node[T] {
	return NewNode(receiver.value, nil, nil)
}

func (receiver *Node[T]) HashCode() string {
	return utils.HashCodeOf(receiver.value)
}

func (receiver *Node[T]) Equals(node *Node[T]) bool {
	return receiver.HashCode() == node.HashCode()
}

func (receiver *Node[T]) LessThan(node *Node[T]) bool {
	return receiver.HashCode() < node.HashCode()
}

func (receiver *Node[T]) ValueEquals(value T) bool {
	return utils.HashCodeOf(receiver.value) == utils.HashCodeOf(value)
}

func (receiver *Node[T]) ValueLessThan(value T) bool {
	return utils.HashCodeOf(receiver.value) < utils.HashCodeOf(value)
}

func (receiver *Node[T]) IsLeaf() bool {
	return receiver.left == nil && receiver.right == nil
}

func (receiver *Node[T]) IsFull() bool {
	return receiver.left != nil && receiver.right != nil
}

func (receiver *Node[T]) Add(value T) {
	if receiver.ValueEquals(value) {
		return
	}

	newNode := NewLeafNode(value)

	if receiver.ValueLessThan(value) {
		if receiver.left == nil {
			receiver.left = newNode
		} else {
			receiver.left.Add(value)
		}
	} else {
		if receiver.right == nil {
			receiver.right = newNode
		} else {
			receiver.right.Add(value)
		}
	}
}

func (receiver *Node[T]) Remove(value T) *Node[T] {
	if receiver == nil {
		return nil
	}

	if receiver.ValueEquals(value) {
		if receiver.IsLeaf() {
			return nil
		}

		if receiver.left == nil {
			return receiver.right
		}

		if receiver.right == nil {
			return receiver.left
		}

		successor := receiver.right.Min()
		receiver.value = successor.value
		receiver.right = receiver.right.Remove(successor.value)
		return receiver
	}

	if receiver.ValueLessThan(value) {
		receiver.left = receiver.left.Remove(value)
	} else {
		receiver.right = receiver.right.Remove(value)
	}

	return receiver
}

func (receiver *Node[T]) Min() *Node[T] {
	if receiver.left == nil {
		return receiver
	}

	return receiver.left.Min()
}

func (receiver *Node[T]) Max() *Node[T] {
	if receiver.right == nil {
		return receiver
	}

	return receiver.right.Max()
}

func (receiver *Node[T]) Find(value T) *Node[T] {
	if receiver == nil {
		return nil
	}

	if receiver.ValueEquals(value) {
		return receiver
	}

	if receiver.ValueLessThan(value) {
		return receiver.left.Find(value)
	}

	return receiver.right.Find(value)
}

func (receiver *Node[T]) Height() int {
	if receiver == nil {
		return 0
	}

	return int(math.Max(float64(receiver.left.Height()), float64(receiver.right.Height()))) + 1
}

func (receiver *Node[T]) ForEach(appliedFunc func(int, T)) {
	if receiver.left != nil {
		receiver.left.ForEach(appliedFunc)
	}
	appliedFunc(-1, receiver.value)
	if receiver.right != nil {
		receiver.right.ForEach(appliedFunc)
	}
}

func (receiver *Node[T]) Has(value T) bool {
	if receiver == nil {
		return false
	}

	if receiver.ValueEquals(value) {
		return true
	}

	if receiver.ValueLessThan(value) {
		return receiver.left.Has(value)
	}

	return receiver.right.Has(value)
}
