package set

import (
	"generic-collections/interfaces"
	"generic-collections/utils"
)

// Set implements the ICollection interface
type Set[T any] struct {
	elements map[string]T
	count    int
}

// New creates a new empty set.
func New[T any]() *Set[T] {
	return &Set[T]{elements: make(map[string]T)}
}

// From creates a new set from a slice of elements.
func From[T any](items ...T) *Set[T] {
	var set = New[T]()
	for _, item := range items {
		var key = utils.HashCodeOf(item)
		set.elements[key] = item
	}
	set.count = len(items)

	return set
}

// ForEach iterates over the elements of the set.
// First argument of the appliedFunc is always 0 because sets do not have indexes.
// Second argument of the appliedFunc is the element of the set.
func (receiver *Set[T]) ForEach(appliedFunc func(int, T)) {
	for _, element := range receiver.elements {
		appliedFunc(0, element)
	}
}

// Add adds an element to the set.
// Overwrites the element if it already exists.
func (receiver *Set[T]) Add(item T) interfaces.ICollection[T] {
	var key = utils.HashCodeOf(item)

	if !receiver.Has(item) {
		receiver.count++
	}

	receiver.elements[key] = item

	return receiver
}

// AddAll adds all elements of the given collection to the set.
// Overwrites the element if it already exists.
func (receiver *Set[T]) AddAll(items interfaces.ICollection[T]) interfaces.ICollection[T] {
	for _, item := range items.ToSlice() {
		receiver.Add(item)
	}

	receiver.count = len(receiver.elements)

	return receiver
}

// Count returns the number of elements in the set.
func (receiver *Set[T]) Count() int {
	return receiver.count
}

// Has checks if the set contains the specified item.
func (receiver *Set[T]) Has(item T) bool {
	var key = utils.HashCodeOf(item)
	var _, ok = receiver.elements[key]
	return ok
}

// HasAll checks if the set contains all the items of the specified collection.
func (receiver *Set[T]) HasAll(items interfaces.ICollection[T]) bool {
	var hasAll = true
	items.ForEach(func(_ int, element T) {
		if !receiver.Has(element) {
			hasAll = false
		}
	})

	return hasAll
}

// Clear removes the specified item from the set.
// Returns the set itself.
func (receiver *Set[T]) Clear() interfaces.ICollection[T] {
	receiver.elements = make(map[string]T)
	receiver.count = 0
	return receiver
}

// Filter removes all elements from the set that do not satisfy the predicate function.
// Returns the set itself.
func (receiver *Set[T]) Filter(predicateFunc func(T) bool) interfaces.ICollection[T] {
	var filtered = New[T]()
	receiver.ForEach(func(index int, element T) {
		if predicateFunc(element) {
			filtered.Add(element)
		}
	})

	return filtered
}

// ToSlice converts the set to a slice.
func (receiver *Set[T]) ToSlice() []T {
	var slice = make([]T, 0)
	for _, element := range receiver.elements {
		slice = append(slice, element)
	}
	return slice
}

// IsEmpty checks if the set is empty.
func (receiver *Set[T]) IsEmpty() bool {
	return receiver.Count() == 0
}

// Clone returns a new set with the same elements.
func (receiver *Set[T]) Clone() interfaces.ICollection[T] {
	var set = New[T]()
	for _, element := range receiver.elements {
		set.Add(element)
	}
	return set
}
