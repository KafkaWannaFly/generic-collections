package set

import (
	"github.com/KafkaWannaFly/generic-collections/hashmap"
	"github.com/KafkaWannaFly/generic-collections/interfaces"
	"github.com/KafkaWannaFly/generic-collections/utils"
)

// Set implements the ICollection interface.
// It represents a collection of unique elements.
type Set[T any] struct {
	elements map[string]T
	count    int
}

var _ interfaces.ICollection[any] = (*Set[any])(nil)

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

// region ICollection[TItem] implementation

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

// HasAny checks if the set contains any of the items of the specified collection.
func (receiver *Set[T]) HasAny(items interfaces.ICollection[T]) bool {
	var hasAny = false
	items.ForEach(func(_ int, element T) {
		if receiver.Has(element) {
			hasAny = true
		}
	})

	return hasAny
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
// Original set is not modified.
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

// Default Return a default empty set
func (receiver *Set[T]) Default() interfaces.ICollection[T] {
	return New[T]()
}

// endregion

// region Set[TItem] specific methods

// Union returns a new set that contains all elements of the set and the specified set.
// Does not modify the original sets.
func (receiver *Set[T]) Union(set *Set[T]) *Set[T] {
	var union = receiver.Clone().(*Set[T])
	union.AddAll(set)
	return union
}

// Intersect returns a new set that contains all elements that are in both the set and the specified set.
// Does not modify the original sets.
func (receiver *Set[T]) Intersect(set *Set[T]) *Set[T] {
	var intersect = New[T]()
	receiver.ForEach(func(_ int, element T) {
		if set.Has(element) {
			intersect.Add(element)
		}
	})
	return intersect
}

// Difference returns a new set that contains all elements that are in the set but not in the specified set.
// Does not modify the original sets.
func (receiver *Set[T]) Difference(set *Set[T]) *Set[T] {
	var difference = New[T]()
	receiver.ForEach(func(_ int, element T) {
		if !set.Has(element) {
			difference.Add(element)
		}
	})
	return difference
}

// SymmetricDifference returns a new set that contains all elements that are in the set or the specified set but not in both.
// Does not modify the original sets.
func (receiver *Set[T]) SymmetricDifference(set *Set[T]) *Set[T] {
	var symmetricDifference = New[T]()
	receiver.ForEach(func(_ int, element T) {
		if !set.Has(element) {
			symmetricDifference.Add(element)
		}
	})

	set.ForEach(func(_ int, element T) {
		if !receiver.Has(element) {
			symmetricDifference.Add(element)
		}
	})

	return symmetricDifference
}

// Map method refers to the Map function.
func (receiver *Set[T]) Map(mapper func(int, T) any) *Set[any] {
	return Map(receiver, mapper)
}

// Reduce method refers to the Reduce function.
func (receiver *Set[T]) Reduce(reducer func(any, T) any, initialValue any) any {
	return Reduce(receiver, reducer, initialValue)
}

// GroupBy method refers to the GroupBy function.
func (receiver *Set[T]) GroupBy(keySelector func(T) any) *hashmap.HashMap[any, *Set[T]] {
	return GroupBy(receiver, keySelector)
}

// endregion

// region Package functions

// IsSet checks if the specified item is a set of type T.
func IsSet[T any](item any) bool {
	if item == nil {
		return false
	}

	_, ok := item.(*Set[T])

	return ok
}

// endregion
