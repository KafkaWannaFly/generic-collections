package stack

import (
	"generic-collections/interfaces"
	"generic-collections/linkedlist"
)

// Stack represents a LIFO (Last In First Out) collection
type Stack[T any] struct {
	*linkedlist.LinkedList[T]
}

var _ interfaces.IIndexableCollection[int, any] = (*Stack[any])(nil)

// New creates a new empty stack
func New[T any]() *Stack[T] {
	return &Stack[T]{LinkedList: linkedlist.New[T]()}
}

// From creates a new stack from a slice of items
func From[T any](items ...T) *Stack[T] {
	return &Stack[T]{LinkedList: linkedlist.From[T](items...)}
}

// region Stack[T] methods

// Push adds an item to the top of the stack
func (receiver *Stack[T]) Push(item T) {
	receiver.AddFirst(item)
}

// Pop removes and returns the item at the top of the stack
// Panics if the stack is empty
func (receiver *Stack[T]) Pop() T {
	return receiver.RemoveFirst()
}

// TryPop removes and returns the item at the top of the stack
// Returns default value and false if the stack is empty
// Otherwise, returns the item and true
func (receiver *Stack[T]) TryPop() (T, bool) {
	return receiver.TryRemoveAt(0)
}

// Peek returns the item at the top of the stack without removing it
// Panics if the stack is empty
func (receiver *Stack[T]) Peek() T {
	return receiver.GetAt(0)
}

// TryPeek returns the item at the top of the stack without removing it
// Returns default value and false if the stack is empty
// Otherwise, returns the item and true
func (receiver *Stack[T]) TryPeek() (T, bool) {
	return receiver.TryGetAt(0)
}

// endregion
