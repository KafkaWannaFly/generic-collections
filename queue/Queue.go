package queue

import (
	"github.com/KafkaWannaFly/generic-collections/hashmap"
	"github.com/KafkaWannaFly/generic-collections/interfaces"
	"github.com/KafkaWannaFly/generic-collections/linkedlist"
)

// Queue represents a FIFO (First In First Out) collection.
type Queue[T any] struct {
	super *linkedlist.LinkedList[T]
}

var _ interfaces.IIndexableCollection[int, any] = (*Queue[any])(nil)

// New creates a new empty queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{super: linkedlist.New[T]()}
}

// From creates a new queue from a slice of items.
func From[T any](items ...T) *Queue[T] {
	return &Queue[T]{super: linkedlist.From[T](items...)}
}

// region interfaces.IIndexableCollection[int, T]

// ForEach iterates over the items in the queue.
func (receiver *Queue[T]) ForEach(apply func(int, T)) {
	receiver.super.ForEach(apply)
}

// Add adds an item to the end of the queue.
func (receiver *Queue[T]) Add(item T) interfaces.ICollection[T] {
	return receiver.super.Add(item)
}

// AddAll adds all items from a collection to the end of the queue.
func (receiver *Queue[T]) AddAll(collection interfaces.ICollection[T]) interfaces.ICollection[T] {
	return receiver.super.AddAll(collection)
}

// Count returns the number of items in the queue.
func (receiver *Queue[T]) Count() int {
	return receiver.super.Count()
}

// Has returns true if the queue contains the specified item.
func (receiver *Queue[T]) Has(item T) bool {
	return receiver.super.Has(item)
}

// HasAll returns true if the queue contains all items from a collection.
func (receiver *Queue[T]) HasAll(collection interfaces.ICollection[T]) bool {
	return receiver.super.HasAll(collection)
}

// HasAny returns true if the queue contains any item from a collection.
func (receiver *Queue[T]) HasAny(collection interfaces.ICollection[T]) bool {
	return receiver.super.HasAny(collection)
}

// Clear removes all items from the queue.
// Returns the queue itself.
func (receiver *Queue[T]) Clear() interfaces.ICollection[T] {
	return receiver.super.Clear()
}

// Filter returns a new queue containing only the items that match the predicate.
// The original queue remains unchanged.
func (receiver *Queue[T]) Filter(predicate func(T) bool) interfaces.ICollection[T] {
	return &Queue[T]{super: receiver.super.Filter(predicate).(*linkedlist.LinkedList[T])}
}

// ToSlice returns a slice containing all items in the queue.
func (receiver *Queue[T]) ToSlice() []T {
	return receiver.super.ToSlice()
}

// IsEmpty returns true if the queue is empty.
func (receiver *Queue[T]) IsEmpty() bool {
	return receiver.super.IsEmpty()
}

// Clone returns a shallow copy of the queue.
func (receiver *Queue[T]) Clone() interfaces.ICollection[T] {
	return &Queue[T]{super: receiver.super.Clone().(*linkedlist.LinkedList[T])}
}

// GetAt returns the item at the specified index.
// Panics if the index is out of range.
func (receiver *Queue[T]) GetAt(index int) T {
	return receiver.super.GetAt(index)
}

// SetAt sets the item at the specified index.
// Panics if the index is out of range.
func (receiver *Queue[T]) SetAt(index int, value T) {
	receiver.super.SetAt(index, value)
}

// TryGetAt returns the item at the specified index.
// Returns default value and false if the index is out of range.
// Otherwise, returns the item and true.
func (receiver *Queue[T]) TryGetAt(index int) (T, bool) {
	return receiver.super.TryGetAt(index)
}

// TrySetAt sets the item at the specified index.
// Returns false if the index is out of range.
// Otherwise, returns true.
func (receiver *Queue[T]) TrySetAt(index int, value T) bool {
	return receiver.super.TrySetAt(index, value)
}

// AddFirst adds an item to the top of the queue.
// Returns the queue itself.
func (receiver *Queue[T]) AddFirst(value T) interfaces.ICollection[T] {
	return receiver.super.AddFirst(value)
}

// AddLast adds an item to the bottom of the queue.
// Returns the queue itself.
func (receiver *Queue[T]) AddLast(value T) interfaces.ICollection[T] {
	return receiver.super.AddLast(value)
}

// AddBefore adds an item before the specified index.
// Returns the queue itself.
// Panics if the index is out of range.
func (receiver *Queue[T]) AddBefore(index int, value T) interfaces.ICollection[T] {
	return receiver.super.AddBefore(index, value)
}

// TryAddBefore adds an item before the specified index.
// Returns false if the index is out of range.
// Otherwise, returns true.
func (receiver *Queue[T]) TryAddBefore(index int, value T) bool {
	return receiver.super.TryAddBefore(index, value)
}

// AddAfter adds an item after the specified index.
// Returns the queue itself.
func (receiver *Queue[T]) AddAfter(index int, value T) interfaces.ICollection[T] {
	return receiver.super.AddAfter(index, value)
}

// TryAddAfter adds an item after the specified index.
// Returns false if the index is out of range.
// Otherwise, returns true.
func (receiver *Queue[T]) TryAddAfter(index int, value T) bool {
	return receiver.super.TryAddAfter(index, value)
}

// RemoveFirst removes and returns the item at the top of the queue.
// Panics if the queue is empty.
func (receiver *Queue[T]) RemoveFirst() T {
	return receiver.super.RemoveFirst()
}

// RemoveLast removes and returns the item at the bottom of the queue.
// Panics if the queue is empty.
func (receiver *Queue[T]) RemoveLast() T {
	return receiver.super.RemoveLast()
}

// RemoveAt removes the specified item from the queue.
// Panics if the index is out of range.
func (receiver *Queue[T]) RemoveAt(index int) T {
	return receiver.super.RemoveAt(index)
}

// TryRemoveAt removes the specified item from the queue.
// Returns default value and false if the index is out of range.
// Otherwise, returns the item and true.
func (receiver *Queue[T]) TryRemoveAt(index int) (T, bool) {
	return receiver.super.TryRemoveAt(index)
}

// FindFirst returns the index of the first item that matches the predicate.
// Returns -1 if no item matches the predicate.
func (receiver *Queue[T]) FindFirst(predicate func(int, T) bool) int {
	return receiver.super.FindFirst(predicate)
}

// FindLast returns the index of the last item that matches the predicate.
// Returns -1 if no item matches the predicate.
func (receiver *Queue[T]) FindLast(predicate func(int, T) bool) int {
	return receiver.super.FindLast(predicate)
}

// FindAll returns the indexes of all items that match the predicate.
func (receiver *Queue[T]) FindAll(predicate func(int, T) bool) []int {
	return receiver.super.FindAll(predicate)
}

// Default returns an empty queue.
func (receiver *Queue[T]) Default() interfaces.ICollection[T] {
	return New[T]()
}

// endregion

// region Queue[T] methods.

// Dequeue removes and returns the item at the beginning of the queue.
// Panics if the queue is empty.
func (receiver *Queue[T]) Dequeue() T {
	return receiver.RemoveFirst()
}

// TryDequeue removes and returns the item at the beginning of the queue.
// Returns default value and false if the queue is empty.
// Otherwise, returns the item and true.
func (receiver *Queue[T]) TryDequeue() (T, bool) {
	return receiver.TryRemoveAt(0)
}

// Enqueue adds an item to the end of the queue.
func (receiver *Queue[T]) Enqueue(item T) {
	receiver.Add(item)
}

// Peek returns the item at the beginning of the queue.
// Panics if the queue is empty.
func (receiver *Queue[T]) Peek() T {
	return receiver.GetAt(0)
}

// TryPeek returns the item at the beginning of the queue.
// Returns default value and false if the queue is empty.
// Otherwise, returns the item and true.
func (receiver *Queue[T]) TryPeek() (T, bool) {
	return receiver.TryGetAt(0)
}

// Map applies a function to each item in the queue and returns a new queue with the results.
// The original queue remains unchanged.
func (receiver *Queue[T]) Map(mapper func(int, T) any) *Queue[any] {
	return Map(receiver, mapper)
}

// Reduce applies a function to each item in the queue and returns the accumulated result.
// The original queue remains unchanged.
func (receiver *Queue[T]) Reduce(reducer func(any, T) any, initial any) any {
	return Reduce(receiver, reducer, initial)
}

// GroupBy groups the items in the queue by the specified key.
// Returns a map where the key is the result of the keySelector function.
// The original queue remains unchanged.
func (receiver *Queue[T]) GroupBy(keySelector func(T) any) *hashmap.HashMap[any, *Queue[T]] {
	return GroupBy(receiver, keySelector)
}

// endregion

// region package methods.

// IsQueue returns true if the collection is a queue.
func IsQueue[T any](collection interfaces.ICollection[T]) bool {
	if collection == nil {
		return false
	}

	_, ok := collection.(*Queue[T])
	return ok
}

// endregion
