package linkedlist

import (
	"github.com/KafkaWannaFly/generic-collections/doctor"
	"github.com/KafkaWannaFly/generic-collections/gc"
	"github.com/KafkaWannaFly/generic-collections/guard"
	"github.com/KafkaWannaFly/generic-collections/hashmap"
	"github.com/KafkaWannaFly/generic-collections/interfaces"
	"github.com/KafkaWannaFly/generic-collections/set"
	"github.com/KafkaWannaFly/generic-collections/utils"
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

var _ interfaces.IIndexableCollection[int, any] = (*LinkedList[any])(nil)

// region ICollection[T]

// ForEach loops through each item in LinkedList and apply function
func (receiver *LinkedList[T]) ForEach(appliedFunc func(int, T)) {
	var curr = receiver.Head
	for i := 0; curr != nil; i++ {
		appliedFunc(i, curr.Value)
		curr = curr.Next
	}
}

// Add an item to the tail of LinkedList.
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
		if utils.HashCodeOf(item) == curr.HashCode() {
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

// endregion

// region IIndexableGetSet[int, T]

// GetAt item with certain index in LinkedList.
// Panic if index out of range or less than 0
func (receiver *LinkedList[T]) GetAt(index int) T {
	guard.EnsureIndexRange(index, receiver.count)

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

// SetAt value to index.
// Panic if index out of range or less than 0
func (receiver *LinkedList[T]) SetAt(index int, value T) {
	guard.EnsureIndexRange(index, receiver.count)

	var curr = receiver.Head
	for i := 0; curr != nil; i++ {
		if index == i {
			curr.Value = value
			break
		}

		curr = curr.Next
	}
}

// TryGetAt item with certain index in LinkedList.
// Return the value and true if index in range, else default value and false
func (receiver *LinkedList[T]) TryGetAt(index int) (T, bool) {
	defer doctor.RecoverDefaultFalse[T]()

	return receiver.GetAt(index), true
}

// TrySetAt value to index.
// Return true if index in range, else false
func (receiver *LinkedList[T]) TrySetAt(index int, value T) bool {
	defer doctor.RecoverFalse()

	receiver.SetAt(index, value)
	return true
}

// endregion

// region IIndexableAdder[int, T]

// AddFirst an item to the head of LinkedList.
// Return LinkedList after modification.
func (receiver *LinkedList[T]) AddFirst(item T) interfaces.ICollection[T] {
	if receiver.Head == nil {
		receiver.Add(item)
	} else {
		var node = NodeOf(item)
		node.Next = receiver.Head
		receiver.Head = node
		receiver.count++
	}

	return receiver
}

// AddLast an item to the tail of LinkedList. Same as Add
// Return LinkedList after modification.
func (receiver *LinkedList[T]) AddLast(item T) interfaces.ICollection[T] {
	return receiver.Add(item)
}

// AddBefore an item before certain index.
// Return LinkedList after modification.
func (receiver *LinkedList[T]) AddBefore(index int, item T) interfaces.ICollection[T] {
	guard.EnsureIndexRange(index, receiver.count+1)

	if index == 0 {
		return receiver.AddFirst(item)
	}

	if index == receiver.count {
		return receiver.AddLast(item)
	}

	var curr = receiver.Head
	for i := 0; curr != nil; i++ {
		if i == index-1 {
			var node = NodeOf(item)
			node.Next = curr.Next
			curr.Next = node
			receiver.count++
			break
		}

		curr = curr.Next
	}

	return receiver
}

// AddAfter an item after certain index.
// Return LinkedList after modification.
func (receiver *LinkedList[T]) AddAfter(index int, item T) interfaces.ICollection[T] {
	guard.EnsureIndexRange(index+1, receiver.count+1)

	if index == -1 {
		return receiver.AddFirst(item)
	}

	if index == receiver.count-1 {
		return receiver.AddLast(item)
	}

	var curr = receiver.Head
	for i := 0; curr != nil; i++ {
		if i == index {
			var node = NodeOf(item)
			node.Next = curr.Next
			curr.Next = node
			receiver.count++
			break
		}

		curr = curr.Next
	}

	return receiver
}

// TryAddBefore an item before certain index.
// Return true if index in range, else false.
func (receiver *LinkedList[T]) TryAddBefore(index int, item T) bool {
	defer doctor.RecoverFalse()

	receiver.AddBefore(index, item)
	return true
}

// TryAddAfter an item after certain index.
// Return true if index in range, else false.
func (receiver *LinkedList[T]) TryAddAfter(index int, item T) bool {
	defer doctor.RecoverFalse()

	receiver.AddAfter(index, item)
	return true
}

// endregion

// region IIndexableRemover[int, T]

// RemoveAt item from LinkedList at certain index.
// Return the removed item.
// Panic if index out of range or less than 0
func (receiver *LinkedList[T]) RemoveAt(index int) T {
	guard.EnsureIndexRange(index, receiver.count)

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
				receiver.Tail = beforeRemovedItem
			}

			break
		}

		curr = curr.Next
	}

	receiver.count--

	return removedItemValue
}

// TryRemoveAt item from LinkedList at certain index.
// Return the removed item and true.
// Return default value and false if index out of range.
func (receiver *LinkedList[T]) TryRemoveAt(index int) (T, bool) {
	defer doctor.RecoverDefaultFalse[T]()

	return receiver.RemoveAt(index), true
}

// RemoveFirst item from LinkedList.
// Return the removed item.
func (receiver *LinkedList[T]) RemoveFirst() T {
	return receiver.RemoveAt(0)
}

// RemoveLast item from LinkedList.
// Return the removed item.
func (receiver *LinkedList[T]) RemoveLast() T {
	return receiver.RemoveAt(receiver.count - 1)
}

// endregion

// region IIndexableFinder[int, T]

// FindFirst item based on predicate.
// Return first matched index if found, else -1
func (receiver *LinkedList[T]) FindFirst(predicate func(int, T) bool) int {
	var isFoundYet = false
	var index = -1
	receiver.ForEach(func(i int, item T) {
		if predicate(i, item) && !isFoundYet {
			index = i
			isFoundYet = true
		}
	})

	return index
}

// FindLast item based on predicate.
// Return last matched index if found, else -1
func (receiver *LinkedList[T]) FindLast(predicate func(int, T) bool) int {
	var isFoundYet = false
	var index = -1
	receiver.ForEach(func(i int, item T) {
		if predicate(i, item) {
			index = i
			isFoundYet = true
		}
	})

	if !isFoundYet {
		return -1
	}

	return index
}

// FindAll items based on predicate.
// Return all matched indexes.
func (receiver *LinkedList[T]) FindAll(predicate func(int, T) bool) []int {
	var indexes = make([]int, 0)
	receiver.ForEach(func(i int, item T) {
		if predicate(i, item) {
			indexes = append(indexes, i)
		}
	})

	return indexes
}

// Default return a default empty LinkedList
func (receiver *LinkedList[T]) Default() interfaces.ICollection[T] {
	return New[T]()
}

// endregion

// region ISlicer[T]

// Slice returns a new collection that contains a slice of the original collection.
// Refer to gc.Slice for more information.
func (receiver *LinkedList[T]) Slice(index int, length int) interfaces.IIndexableCollection[int, T] {
	return gc.Slice[T](receiver, index, length)
}

// endregion

// region LinkedList[T] specific methods

// NodeAt get Node object at certain index.
func (receiver *LinkedList[T]) NodeAt(index int) *Node[T] {
	guard.EnsureIndexRange(index, receiver.count)

	curr := receiver.Head
	for i := 0; curr != nil; i++ {
		if i == index {
			return curr
		}

		curr = curr.Next
	}

	return nil
}

// TryNodeAt get Node object at certain index.
// Return the Node object and true if index in range, else nil and false.
func (receiver *LinkedList[T]) TryNodeAt(index int) (*Node[T], bool) {
	defer doctor.RecoverDefaultFalse[*Node[T]]()

	return receiver.NodeAt(index), true
}

// Map applies the given mapper function to each element of the list.
// Returns a new list containing the results. Don't modify the original list.
func (receiver *LinkedList[T]) Map(mapper func(int, T) any) *LinkedList[any] {
	return Map(receiver, mapper)
}

// Reduce applies the given reducer function to each element of the list.
// Returns the accumulated result.
func (receiver *LinkedList[T]) Reduce(reducer func(any, T) any, initialValue any) any {
	return Reduce(receiver, reducer, initialValue)
}

// GroupBy groups the elements of the list by the specified key.
// Returns a map where the key is the result of the keySelector function
func (receiver *LinkedList[T]) GroupBy(keySelector func(T) any) *hashmap.HashMap[any, *LinkedList[T]] {
	return GroupBy(receiver, keySelector)
}

// endregion

// region package methods

// IsLinkedList check if an object is a LinkedList of type T
func IsLinkedList[T any](item any) bool {
	if item == nil {
		return false
	}

	_, ok := item.(*LinkedList[T])
	return ok
}

//
