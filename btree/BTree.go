package btree

import "github.com/KafkaWannaFly/generic-collections/interfaces"

type BTree[T any] struct {
	root  *Node[T]
	count int
}

var _ interfaces.ICollection[int] = (*BTree[int])(nil)

func New[T any]() *BTree[T] {
	return &BTree[T]{}
}

func From[T any](elements ...T) *BTree[T] {
	tree := New[T]()
	for _, element := range elements {
		tree.Add(element)
	}
	return tree
}

func (receiver *BTree[T]) Add(item T) interfaces.ICollection[T] {
	if receiver.root == nil {
		receiver.root = NewLeafNode(item)
	} else {
		receiver.root.Add(item)
	}
	receiver.count++
	return receiver
}

func (receiver *BTree[T]) AddAll(i interfaces.ICollection[T]) interfaces.ICollection[T] {
	for _, item := range i.ToSlice() {
		receiver.Add(item)
	}
	return receiver
}

func (receiver *BTree[T]) Count() int {
	return receiver.count
}

func (receiver *BTree[T]) ForEach(appliedFunc func(int, T)) {
	if receiver.root == nil {
		return
	}
	receiver.root.ForEach(appliedFunc)
}

func (receiver *BTree[T]) Has(item T) bool {
	if receiver.root == nil {
		return false
	}

	return receiver.root.Has(item)
}

func (receiver *BTree[T]) HasAll(i interfaces.ICollection[T]) bool {
	if receiver.root == nil {
		return false
	}

	for _, item := range i.ToSlice() {
		if !receiver.Has(item) {
			return false
		}
	}

	return true
}

func (receiver *BTree[T]) HasAny(i interfaces.ICollection[T]) bool {
	if receiver.root == nil {
		return false
	}

	for _, item := range i.ToSlice() {
		if receiver.Has(item) {
			return true
		}
	}
	return false
}

func (receiver *BTree[T]) Remove(item T) interfaces.ICollection[T] {
	if receiver.root == nil {
		return receiver
	}

	receiver.root.Remove(item)
	receiver.count--
	return receiver
}

func (receiver *BTree[T]) Height() int {
	if receiver.root == nil {
		return 0
	}
	return receiver.root.Height()
}

func (receiver *BTree[T]) Clear() interfaces.ICollection[T] {
	receiver.root = nil
	receiver.count = 0
	return receiver
}

func (receiver *BTree[T]) Filter(predicate func(T) bool) interfaces.ICollection[T] {
	var filtered = New[T]()

	receiver.ForEach(func(index int, item T) {
		if predicate(item) {
			filtered.Add(item)
		}
	})

	return filtered
}

func (receiver *BTree[T]) ToSlice() []T {
	var slice = make([]T, 0, receiver.count)
	receiver.ForEach(func(index int, item T) {
		slice = append(slice, item)
	})
	return slice
}

func (receiver *BTree[T]) IsEmpty() bool {
	return receiver.count == 0
}

func (receiver *BTree[T]) Clone() interfaces.ICollection[T] {
	var clone = New[T]()
	receiver.ForEach(func(index int, item T) {
		clone.Add(item)
	})

	return clone
}

func (receiver *BTree[T]) Default() interfaces.ICollection[T] {
	return New[T]()
}
