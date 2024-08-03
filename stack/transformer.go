package stack

import "generic-collections/hashmap"

// Map creates a new stack by applying a mapper function to each item in the original stack
// The original stack remains unchanged
func Map[TType any, TResult any](collection *Stack[TType], mapper func(TType) TResult) *Stack[TResult] {
	result := New[TResult]()

	collection.ForEach(func(i int, item TType) {
		result.AddLast(mapper(item))
	})

	return result
}

// Reduce reduces the stack to a single value by accumulating items
// The accumulator function receives the current accumulated value and the next item
func Reduce[TType any, TResult any](collection *Stack[TType], reducer func(TResult, TType) TResult, initialValue TResult) TResult {
	var result = initialValue
	collection.ForEach(func(_ int, item TType) {
		result = reducer(result, item)
	})

	return result
}

// GroupBy groups the elements of the stack by the specified key
// Returns a map where the key is the result of the keySelector function
func GroupBy[TType any, TKey any](collection *Stack[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *Stack[TType]] {
	groups := hashmap.New[TKey, *Stack[TType]]()

	collection.ForEach(func(i int, item TType) {
		key := keySelector(item)
		if !groups.HasKey(key) {
			groups.Put(key, New[TType]())
		}

		groups.Get(key).Push(item)
	})

	return groups
}
