package stack

import (
	"generic-collections/hashmap"
	"generic-collections/interfaces"
	"generic-collections/linkedlist"
)

// Stack represents a LIFO (Last In First Out) collection.
type Stack[T any] struct {
	super *linkedlist.LinkedList[T]
}

var _ interfaces.IIndexableCollection[int, any] = (*Stack[any])(nil)

// New creates a new empty stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{super: linkedlist.New[T]()}
}

// From creates a new stack from a slice of items.
func From[T any](items ...T) *Stack[T] {
	return &Stack[T]{super: linkedlist.From[T](items...)}
}

// region interfaces.IIndexableCollection[int, T]

// ForEach iterates over the items in the stack.
func (receiver *Stack[T]) ForEach(apply func(int, T)) {
	receiver.super.ForEach(apply)
}

// Add adds an item to the bottom stack.
func (receiver *Stack[T]) Add(item T) interfaces.ICollection[T] {
	return receiver.super.Add(item)
}

// AddAll adds all items from a collection to the bottom of the stack.
func (receiver *Stack[T]) AddAll(collection interfaces.ICollection[T]) interfaces.ICollection[T] {
	return receiver.super.AddAll(collection)
}

// Count returns the number of items in the stack.
func (receiver *Stack[T]) Count() int {
	return receiver.super.Count()
}

// Has returns true if the stack contains the specified item.
func (receiver *Stack[T]) Has(item T) bool {
	return receiver.super.Has(item)
}

// HasAll returns true if the stack contains all items from a collection.
func (receiver *Stack[T]) HasAll(collection interfaces.ICollection[T]) bool {
	return receiver.super.HasAll(collection)
}

// HasAny returns true if the stack contains any item from a collection.
func (receiver *Stack[T]) HasAny(collection interfaces.ICollection[T]) bool {
	return receiver.super.HasAny(collection)
}

// Clear removes all items from the stack.
func (receiver *Stack[T]) Clear() interfaces.ICollection[T] {
	return receiver.super.Clear()
}

// Filter returns a new stack containing only the items that match the predicate.
// The original stack remains unchanged.
func (receiver *Stack[T]) Filter(predicate func(T) bool) interfaces.ICollection[T] {
	return &Stack[T]{super: receiver.super.Filter(predicate).(*linkedlist.LinkedList[T])}
}

// ToSlice returns a slice containing all items in the stack.
func (receiver *Stack[T]) ToSlice() []T {
	return receiver.super.ToSlice()
}

// IsEmpty returns true if the stack is empty.
func (receiver *Stack[T]) IsEmpty() bool {
	return receiver.super.IsEmpty()
}

// Clone returns a shallow copy of the stack.
func (receiver *Stack[T]) Clone() interfaces.ICollection[T] {
	return &Stack[T]{super: receiver.super.Clone().(*linkedlist.LinkedList[T])}
}

// GetAt returns the item at the specified index.
// Panics if the index is out of range.
func (receiver *Stack[T]) GetAt(index int) T {
	return receiver.super.GetAt(index)
}

// SetAt sets the item at the specified index.
// Panics if the index is out of range.
func (receiver *Stack[T]) SetAt(index int, value T) {
	receiver.super.SetAt(index, value)
}

// TryGetAt returns the item at the specified index.
// Returns default value and false if the index is out of range.
// Otherwise, returns the item and true.
func (receiver *Stack[T]) TryGetAt(index int) (T, bool) {
	return receiver.super.TryGetAt(index)
}

// TrySetAt sets the item at the specified index.
// Returns false if the index is out of range.
// Otherwise, returns true.
func (receiver *Stack[T]) TrySetAt(index int, value T) bool {
	return receiver.super.TrySetAt(index, value)
}

// AddFirst adds an item to the top of the stack.
// Returns the stack itself.
func (receiver *Stack[T]) AddFirst(value T) interfaces.ICollection[T] {
	return receiver.super.AddFirst(value)
}

// AddLast adds an item to the bottom of the stack.
// Returns the stack itself.
func (receiver *Stack[T]) AddLast(value T) interfaces.ICollection[T] {
	return receiver.super.AddLast(value)
}

// AddBefore adds an item before the specified index.
// Returns the stack itself.
// Panics if the index is out of range.
func (receiver *Stack[T]) AddBefore(index int, value T) interfaces.ICollection[T] {
	return receiver.super.AddBefore(index, value)
}

// TryAddBefore adds an item before the specified index.
// Returns false if the index is out of range.
// Otherwise, returns true.
func (receiver *Stack[T]) TryAddBefore(index int, value T) bool {
	return receiver.super.TryAddBefore(index, value)
}

// AddAfter adds an item after the specified index.
// Returns the stack itself.
func (receiver *Stack[T]) AddAfter(index int, value T) interfaces.ICollection[T] {
	return receiver.super.AddAfter(index, value)
}

// TryAddAfter adds an item after the specified index.
// Returns false if the index is out of range.
// Otherwise, returns true.
func (receiver *Stack[T]) TryAddAfter(index int, value T) bool {
	return receiver.super.TryAddAfter(index, value)
}

// RemoveFirst removes and returns the item at the top of the stack.
// Panics if the stack is empty.
func (receiver *Stack[T]) RemoveFirst() T {
	return receiver.super.RemoveFirst()
}

// RemoveLast removes and returns the item at the bottom of the stack.
// Panics if the stack is empty.
func (receiver *Stack[T]) RemoveLast() T {
	return receiver.super.RemoveLast()
}

// RemoveAt removes the specified item from the stack.
// Panics if the index is out of range.
func (receiver *Stack[T]) RemoveAt(index int) T {
	return receiver.super.RemoveAt(index)
}

// TryRemoveAt removes the specified item from the stack.
// Returns default value and false if the index is out of range.
// Otherwise, returns the item and true.
func (receiver *Stack[T]) TryRemoveAt(index int) (T, bool) {
	return receiver.super.TryRemoveAt(index)
}

// FindFirst returns the index of the first item that matches the predicate.
// Returns -1 if no item matches the predicate.
func (receiver *Stack[T]) FindFirst(predicate func(int, T) bool) int {
	return receiver.super.FindFirst(predicate)
}

// FindLast returns the index of the last item that matches the predicate.
// Returns -1 if no item matches the predicate.
func (receiver *Stack[T]) FindLast(predicate func(int, T) bool) int {
	return receiver.super.FindLast(predicate)
}

// FindAll returns the indexes of all items that match the predicate.
func (receiver *Stack[T]) FindAll(predicate func(int, T) bool) []int {
	return receiver.super.FindAll(predicate)
}

// endregion

// region Stack[T] methods.

// Push adds an item to the top of the stack.
func (receiver *Stack[T]) Push(item T) {
	receiver.super.AddFirst(item)
}

// Pop removes and returns the item at the top of the stack.
// Panics if the stack is empty.
func (receiver *Stack[T]) Pop() T {
	return receiver.super.RemoveFirst()
}

// TryPop removes and returns the item at the top of the stack.
// Returns default value and false if the stack is empty.
// Otherwise, returns the item and true.
func (receiver *Stack[T]) TryPop() (T, bool) {
	return receiver.super.TryRemoveAt(0)
}

// Peek returns the item at the top of the stack without removing it.
// Panics if the stack is empty.
func (receiver *Stack[T]) Peek() T {
	return receiver.super.GetAt(0)
}

// TryPeek returns the item at the top of the stack without removing it.
// Returns default value and false if the stack is empty.
// Otherwise, returns the item and true.
func (receiver *Stack[T]) TryPeek() (T, bool) {
	return receiver.super.TryGetAt(0)
}

// Map creates a new stack by applying a mapper function to each item in the original stack.
// The original stack remains unchanged.
func (receiver *Stack[T]) Map(mapper func(T) any) *Stack[any] {
	return Map(receiver, mapper)
}

// Reduce reduces the stack to a single value by applying an accumulator function to each item.
func (receiver *Stack[T]) Reduce(accumulator func(any, T) any, initialValue any) any {
	return Reduce(receiver, accumulator, initialValue)
}

// GroupBy groups items in the stack by a key returned by a keySelector function.
func (receiver *Stack[T]) GroupBy(keySelector func(T) any) *hashmap.HashMap[any, *Stack[T]] {
	return GroupBy(receiver, keySelector)
}

// endregion.

// region Package functions

// IsStack checks if the specified collection is a stack.
func IsStack[T any](collection any) bool {
	if collection == nil {
		return false
	}

	_, ok := collection.(*Stack[T])

	return ok
}

// endregion
