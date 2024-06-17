package set

import (
	"generic-collections/interfaces"
	"generic-collections/utils"
)

// Set Implement the ICollection interface
type Set[T any] struct {
	elements map[string]T
	count    int
}

func New[T any]() *Set[T] {
	return &Set[T]{elements: make(map[string]T)}
}

func From[T any](items ...T) *Set[T] {
	var set = New[T]()
	for _, item := range items {
		var key = utils.HashCodeOf(item)
		set.elements[key] = item
	}
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

	receiver.elements[key] = item

	return receiver
}

// AddAll adds all elements of the given collection to the set.
// Overwrites the element if it already exists.
func (receiver *Set[T]) AddAll(items interfaces.ICollection[T]) interfaces.ICollection[T] {
	for _, item := range items.ToSlice() {
		if receiver.Has(item) {
			receiver.count++
		}

		receiver.Add(item)
	}

	return receiver
}

func (receiver *Set[T]) Count() int {
	return receiver.count
}

func (receiver *Set[T]) Has(item T) bool {
	var key = utils.HashCodeOf(item)
	var _, ok = receiver.elements[key]
	return ok
}

func (receiver *Set[T]) HasAll(i interfaces.ICollection[T]) bool {
	for _, item := range i.ToSlice() {
		if !receiver.Has(item) {
			return false
		}
	}

	return true
}

func (receiver *Set[T]) Clear() interfaces.ICollection[T] {
	receiver.elements = make(map[string]T)
	receiver.count = 0
	return receiver
}

func (receiver *Set[T]) Filter(predicateFunc func(T) bool) interfaces.ICollection[T] {
	receiver.ForEach(func(index int, element T) {
		if !predicateFunc(element) {
			delete(receiver.elements, utils.HashCodeOf(element))
		}
	})

	return receiver
}

func (receiver *Set[T]) ToSlice() []T {
	var slice = make([]T, 0)
	for _, element := range receiver.elements {
		slice = append(slice, element)
	}
	return slice
}

func (receiver *Set[T]) IsEmpty() bool {
	return receiver.Count() == 0
}
