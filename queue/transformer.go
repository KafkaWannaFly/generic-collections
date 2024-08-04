package queue

import "generic-collections/hashmap"

// Map applies a function to each item in the queue and returns a new queue with the results.
func Map[TType any, TResult any](queue *Queue[TType], mapper func(int, TType) TResult) *Queue[TResult] {
	result := New[TResult]()
	queue.ForEach(func(index int, item TType) {
		result.Add(mapper(index, item))
	})

	return result
}

// Reduce applies a function to each item in the queue and returns the accumulated result.
func Reduce[TType any, TResult any](queue *Queue[TType], reducer func(TResult, TType) TResult, initial TResult) TResult {
	result := initial
	queue.ForEach(func(_ int, item TType) {
		result = reducer(result, item)
	})

	return result
}

// GroupBy groups the items in the queue by the specified key.
func GroupBy[TType any, TKey any](queue *Queue[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *Queue[TType]] {
	groups := hashmap.New[TKey, *Queue[TType]]()

	queue.ForEach(func(_ int, item TType) {
		key := keySelector(item)
		if !groups.HasKey(key) {
			groups.Put(key, From[TType](item))
		} else {
			groups.Get(key).Add(item)
		}
	})

	return groups
}
