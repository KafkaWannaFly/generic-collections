package list

import (
	"github.com/KafkaWannaFly/generic-collections/doctor"
	"github.com/KafkaWannaFly/generic-collections/gc"
	"github.com/KafkaWannaFly/generic-collections/guard"
	"github.com/KafkaWannaFly/generic-collections/hashmap"
	"github.com/KafkaWannaFly/generic-collections/interfaces"
	"github.com/KafkaWannaFly/generic-collections/utils"
)

type List[T any] struct {
	elements []T
	count    int
}

var _ interfaces.IIndexableCollection[int, int] = (*List[int])(nil)

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

// region ICollection[TItem] implementation

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

// HasAny checks if the list contains any of the items of the specified collection.
func (receiver *List[T]) HasAny(items interfaces.ICollection[T]) bool {
	var result = false
	items.ForEach(func(index int, item T) {
		if receiver.Has(item) {
			result = true
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
	var newList = New[T]()
	newList.AddAll(receiver)

	return newList
}

// endregion

// region IIndexableGetSet[TItem] implementation

// GetAt the value of the element at the specified index.
// Panics if the index is out of range.
func (receiver *List[T]) GetAt(i int) T {
	guard.EnsureIndexRange(i, receiver.count)

	return receiver.elements[i]
}

// SetAt the value of the element at the specified index.
// Panics if the index is out of range.
func (receiver *List[T]) SetAt(i int, item T) {
	guard.EnsureIndexRange(i, receiver.count)

	receiver.elements[i] = item
}

// TryGetAt the value of the element at the specified index.
// Returns the value and true if the index is in range, otherwise the default value and false.
func (receiver *List[T]) TryGetAt(i int) (T, bool) {
	defer doctor.RecoverDefaultFalse[T]()

	return receiver.GetAt(i), true
}

// TrySetAt the value of the element at the specified index.
// Returns true if the index is in range, otherwise false.
func (receiver *List[T]) TrySetAt(i int, item T) bool {
	defer doctor.RecoverFalse()

	receiver.SetAt(i, item)
	return true
}

// endregion

// region IIndexableAdder[TItem] implementation

// AddFirst adds the item to the beginning of the list.
// Returns the list itself.
func (receiver *List[T]) AddFirst(item T) interfaces.ICollection[T] {
	receiver.elements = append([]T{item}, receiver.elements...)
	receiver.count++

	return receiver
}

// AddLast adds the item to the end of the list.
// Returns the list itself.
func (receiver *List[T]) AddLast(item T) interfaces.ICollection[T] {
	return receiver.Add(item)
}

// AddBefore adds the item before the element at the specified index.
// Returns the list itself.
func (receiver *List[T]) AddBefore(i int, item T) interfaces.ICollection[T] {
	guard.EnsureIndexRange(i, receiver.count+1)

	receiver.elements = append(receiver.elements[:i], append([]T{item}, receiver.elements[i:]...)...)
	receiver.count++

	return receiver
}

// TryAddBefore adds the item before the element at the specified index.
// Returns true if the index is in range, otherwise false.
func (receiver *List[T]) TryAddBefore(i int, item T) bool {
	defer doctor.RecoverFalse()

	receiver.AddBefore(i, item)
	return true
}

// AddAfter adds the item after the element at the specified index.
// Returns the list itself.
func (receiver *List[T]) AddAfter(i int, item T) interfaces.ICollection[T] {
	guard.EnsureIndexRange(i+1, receiver.count+1)

	receiver.elements = append(receiver.elements[:i+1], append([]T{item}, receiver.elements[i+1:]...)...)
	receiver.count++

	return receiver
}

// TryAddAfter adds the item after the element at the specified index.
// Returns true if the index is in range, otherwise false.
func (receiver *List[T]) TryAddAfter(i int, item T) bool {
	defer doctor.RecoverFalse()

	receiver.AddAfter(i, item)
	return true
}

// endregion

// region IIndexableRemover[TItem] implementation

// RemoveAt item at the specified index.
// Panics if the index is out of range.
func (receiver *List[T]) RemoveAt(i int) T {
	guard.EnsureIndexRange(i, receiver.count)

	var item = receiver.elements[i]
	receiver.elements = append(receiver.elements[:i], receiver.elements[i+1:]...)
	receiver.count--

	return item
}

// TryRemoveAt item at the specified index.
// Returns the value and true if the index is in range, otherwise the default value and false.
func (receiver *List[T]) TryRemoveAt(i int) (T, bool) {
	defer doctor.RecoverDefaultFalse[T]()

	return receiver.RemoveAt(i), true
}

// RemoveFirst item from the beginning of the list.
// Panics if the list is empty.
func (receiver *List[T]) RemoveFirst() T {
	return receiver.RemoveAt(0)
}

// RemoveLast item from the end of the list.
// Panics if the list is empty.
func (receiver *List[T]) RemoveLast() T {
	return receiver.RemoveAt(receiver.count - 1)
}

// endregion

// region IIndexableFinder[TItem] implementation

// FindFirst the first element that satisfies the predicate.
// Returns the index of the element if found, otherwise -1.
func (receiver *List[T]) FindFirst(predicate func(int, T) bool) int {
	var index = -1

	for i, element := range receiver.elements {
		if predicate(i, element) {
			index = i
			break
		}
	}

	return index
}

// FindLast the last element that satisfies the predicate.
// Returns the index of the element if found, otherwise -1.
func (receiver *List[T]) FindLast(predicate func(int, T) bool) int {
	var index = -1

	for i := range receiver.elements {
		reverseIdx := receiver.count - i - 1
		if predicate(reverseIdx, receiver.elements[reverseIdx]) {
			index = reverseIdx
			break
		}
	}

	return index
}

// FindAll items based on predicate.
// Return all matched indexes.
func (receiver *List[T]) FindAll(predicate func(int, T) bool) []int {
	var indexes = make([]int, 0)
	receiver.ForEach(func(i int, item T) {
		if predicate(i, item) {
			indexes = append(indexes, i)
		}
	})

	return indexes
}

// Default returns a new empty list.
func (receiver *List[T]) Default() interfaces.ICollection[T] {
	return New[T]()
}

// region ISlicer[TItem] implementation

// Slice returns a new list containing the elements from the specified index and a specified length.
// Refer to gc.Slice for more information.
func (receiver *List[T]) Slice(index int, length int) interfaces.IIndexableCollection[int, T] {
	return gc.Slice[T](receiver, index, length)
}

// endregion

// region List specific methods

// Map refers to the Map function in Transformers.go.
func (receiver *List[T]) Map(mapper func(int, T) any) *List[any] {
	return Map(receiver, mapper)
}

// Reduce refers to the Reduce function in Transformers.go.
func (receiver *List[T]) Reduce(reducer func(any, T) any, initialValue any) any {
	return Reduce(receiver, reducer, initialValue)
}

// GroupBy refers to the GroupBy function in Transformers.go.
func (receiver *List[T]) GroupBy(keySelector func(T) any) *hashmap.HashMap[any, *List[T]] {
	return GroupBy(receiver, keySelector)
}

// endregion

// region Package functions

// IsList checks if the given collection is a list.
func IsList[T any](collection any) bool {
	if collection == nil {
		return false
	}

	_, ok := collection.(*List[T])

	return ok
}

// endregion
