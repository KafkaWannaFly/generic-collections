package list

import (
	"generic-collections/interfaces"
	"generic-collections/utils"
)

type List[T any] struct {
	elements []T
	count    int
}

// New creates a new empty list.
func New[T any]() *List[T] {
	return &List[T]{elements: make([]T, 0)}
}

// From creates a new list from a slice of elements.
func From[T any](elements ...T) *List[T] {
	var list = &List[T]{elements: elements}
	list.count = len(elements)
	return list
}

// region ICollection[T] implementation

// ForEach iterates over the elements of the list.
// First argument of the appliedFunc is the index of the element.
// Second argument of the appliedFunc is the element of the list.
func (receiver *List[T]) ForEach(appliedFunc func(int, T)) {
	for i, element := range receiver.elements {
		appliedFunc(i, element)
	}
}

// Add new element to the list.
// Returns the list itself.
func (receiver *List[T]) Add(item T) interfaces.ICollection[T] {
	receiver.elements = append(receiver.elements, item)
	receiver.count++

	return receiver
}

// AddAll adds all elements of the given collection to the list.
// Returns the list itself.
func (receiver *List[T]) AddAll(items interfaces.ICollection[T]) interfaces.ICollection[T] {
	receiver.elements = append(receiver.elements, items.ToSlice()...)
	receiver.count = len(receiver.elements)
	return receiver
}

// Count returns the number of elements in the list.
func (receiver *List[T]) Count() int {
	return receiver.count
}

// Has checks if the list contains the specified item.
func (receiver *List[T]) Has(item T) bool {
	for _, element := range receiver.elements {
		if utils.IsEqual(element, item) {
			return true
		}
	}

	return false
}

// HasAll checks if the list contains all the items of the specified collection.
func (receiver *List[T]) HasAll(items interfaces.ICollection[T]) bool {
	elementMap := make(map[string]bool)
	receiver.ForEach(func(index int, element T) {
		var key = utils.HashCodeOf(element)
		elementMap[key] = true
	})

	var result = true
	items.ForEach(func(index int, item T) {
		var key = utils.HashCodeOf(item)
		if _, exists := elementMap[key]; !exists {
			result = false
		}
	})

	return result
}

// Clear removes all elements from the list.
func (receiver *List[T]) Clear() interfaces.ICollection[T] {
	receiver.elements = make([]T, 0)
	receiver.count = 0

	return receiver
}

// Filter returns a new list containing only the elements that satisfy the predicate.
// The original list remains unchanged.
func (receiver *List[T]) Filter(predicate func(T) bool) interfaces.ICollection[T] {
	ans := New[T]()

	for _, element := range receiver.elements {
		if predicate(element) {
			ans.Add(element)
		}
	}

	return ans
}

// ToSlice returns the elements of the list as a slice.
func (receiver *List[T]) ToSlice() []T {
	return receiver.elements
}

// IsEmpty checks if the list is empty.
func (receiver *List[T]) IsEmpty() bool {
	return receiver.count == 0
}

// Clone returns a new list with the same elements.
func (receiver *List[T]) Clone() interfaces.ICollection[T] {
	return From[T](receiver.elements...)
}

// endregion

// region IIndexable[T] implementation

// Get the value of the element at the specified index.
// Panics if the index is out of range.
func (receiver *List[T]) Get(i int) T {
	return receiver.elements[i]
}

// Set the value of the element at the specified index.
// Panics if the index is out of range.
func (receiver *List[T]) Set(i int, item T) {
	receiver.elements[i] = item
}

// Find the first element that satisfies the predicate.
// Returns the index of the element if found, otherwise -1.
func (receiver *List[T]) Find(predicate func(T) bool) int {
	var index = -1

	for i, element := range receiver.elements {
		if predicate(element) {
			index = i
			break
		}
	}

	return index
}

// Remove item at the specified index.
// Panics if the index is out of range.
func (receiver *List[T]) Remove(i int) T {
	var item = receiver.elements[i]
	receiver.elements = append(receiver.elements[:i], receiver.elements[i+1:]...)
	receiver.count--

	return item
}

// endregion
