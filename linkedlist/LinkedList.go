package linkedlist

import (
	"generic-collections/interfaces"
	"generic-collections/set"
	"generic-collections/utils"
)

type LinkedList[T any] struct {
	Head  *Node[T]
	Tail  *Node[T]
	count int
}

// New creates a new empty LinkedList.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// From create a LinkedList from a slice
func From[T any](elements ...T) *LinkedList[T] {
	var list = &LinkedList[T]{}
	for _, element := range elements {
		list.Add(element)
	}
	return list
}

var _ interfaces.ICollection[any] = (*LinkedList[any])(nil)

// ForEach loops through each item in LinkedList and apply function
func (receiver *LinkedList[T]) ForEach(appliedFunc func(int, T)) {
	var curr = receiver.Head
	for i := 0; curr != nil; i++ {
		appliedFunc(i, curr.Value)
		curr = curr.Next
	}
}

// Add an item to LinkedList.
// Return LinkedList after modification
func (receiver *LinkedList[T]) Add(item T) interfaces.ICollection[T] {
	var node = NodeOf(item)
	if receiver.Head == nil {
		receiver.Head = node
		receiver.Tail = node
	} else {
		receiver.Tail.Next = node
		receiver.Tail = node
	}
	receiver.count++
	return receiver
}

// AddAll items from collection to LinkedList.
// Return LinkedList after modification
func (receiver *LinkedList[T]) AddAll(collection interfaces.ICollection[T]) interfaces.ICollection[T] {
	collection.ForEach(func(_ int, item T) {
		receiver.Add(item)
	})

	return receiver
}

// Count number of items inside LinkedList
func (receiver *LinkedList[T]) Count() int {
	return receiver.count
}

// Has check if item exist in LinkedList
func (receiver *LinkedList[T]) Has(item T) bool {
	var curr = receiver.Head

	for curr != nil {
		if utils.HashCodeOf(item) == curr.GetHashCode() {
			return true
		}
		curr = curr.Next
	}

	return false
}

// HasAll check if LinkedList contains all items in collection
func (receiver *LinkedList[T]) HasAll(collection interfaces.ICollection[T]) bool {
	var hasAll = true
	var itemSet = set.From(receiver.ToSlice()...)
	collection.ForEach(func(_ int, item T) {
		if !itemSet.Has(item) {
			hasAll = false
		}
	})

	return hasAll
}

// HasAny check if LinkedList contains at least 1 item from collection
func (receiver *LinkedList[T]) HasAny(collection interfaces.ICollection[T]) bool {
	var hasAny = false
	var itemSet = set.From(receiver.ToSlice()...)
	collection.ForEach(func(_ int, item T) {
		if itemSet.Has(item) {
			hasAny = true
		}
	})

	return hasAny
}

// Clear all LinkedList.
// Return LinkedList after modification
func (receiver *LinkedList[T]) Clear() interfaces.ICollection[T] {
	receiver.Head = nil
	receiver.Tail = nil
	receiver.count = 0

	return receiver
}

// Filter LinkedList based on predicate function. It doesn't modify current LinkedList.
// Return a new LinkedList after filtering
func (receiver *LinkedList[T]) Filter(predicate func(T) bool) interfaces.ICollection[T] {
	var linkedList = New[T]()

	receiver.ForEach(func(_ int, item T) {
		if predicate(item) {
			linkedList.Add(item)
		}
	})

	return linkedList
}

// ToSlice convert LinkedList to a slice
func (receiver *LinkedList[T]) ToSlice() []T {
	var slice = make([]T, receiver.count)

	receiver.ForEach(func(i int, item T) {
		slice[i] = item
	})

	return slice
}

// IsEmpty check if LinkedList has any item
func (receiver *LinkedList[T]) IsEmpty() bool {
	return receiver.count == 0
}

// Clone create a copy of current LinkedList.
// Return a copy of LinkedList
func (receiver *LinkedList[T]) Clone() interfaces.ICollection[T] {
	var linkedList = New[T]()
	linkedList.AddAll(receiver)

	return linkedList
}

// Get item with certain index in LinkedList.
// Panic if index out of range or less than 0
func (receiver *LinkedList[T]) Get(index int) T {
	if index > receiver.count-1 || index < 0 {
		panic("Index out of range")
	}

	var curr = receiver.Head
	var val T
	for i := 0; curr != nil; i++ {
		if i == index {
			val = curr.Value
		}

		curr = curr.Next
	}

	return val
}

// Set value to index.
// Panic if index out of range or less than 0
func (receiver *LinkedList[T]) Set(index int, value T) {
	if index > receiver.count-1 || index < 0 {
		panic("Index out of range")
	}

	var curr = receiver.Head
	for i := 0; curr != nil; i++ {
		if index == i {
			curr.Value = value
			break
		}

		curr = curr.Next
	}
}

// Find item based on predicate.
// Return first matched index if found, else -1
func (receiver *LinkedList[T]) Find(predicate func(T) bool) int {
	var isFoundYet = false
	var index = -1
	receiver.ForEach(func(i int, item T) {
		if predicate(item) && !isFoundYet {
			index = i
			isFoundYet = true
		}
	})

	return index
}

// Remove item from LinkedList at certain index.
// Return the removed item.
// Panic if index out of range or less than 0
func (receiver *LinkedList[T]) Remove(index int) T {
	if index > receiver.count-1 || index < 0 {
		panic("Index out of range")
	}

	var curr = receiver.Head
	var removedItemValue T
	for i := 0; curr != nil; i++ {
		if index == 0 {
			// If remove the head
			removedItemValue = receiver.Head.Value
			receiver.Head = receiver.Head.Next
			break
		}

		if index-1 == i {
			// If remove at middle of the list
			var beforeRemovedItem = curr
			var tobeRemovedItem = beforeRemovedItem.Next
			var afterRemovedItem = tobeRemovedItem.Next
			removedItemValue = tobeRemovedItem.Value

			beforeRemovedItem.Next = afterRemovedItem

			if index == receiver.count-1 {
				// If remove at the tail
				receiver.Tail = afterRemovedItem
			}

			break
		}

		curr = curr.Next
	}

	return removedItemValue
}
