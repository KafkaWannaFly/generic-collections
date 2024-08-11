package list

import (
	"github.com/KafkaWannaFly/generic-collections/hashmap"
)

// Map applies the given mapper function to each element of the list.
// Returns a new list containing the results.
func Map[TType any, TResult any](list *List[TType], mapper func(int, TType) TResult) *List[TResult] {
	var result = New[TResult]()
	list.ForEach(func(index int, item TType) {
		result.Add(mapper(index, item))
	})
	return result
}

// Reduce applies the given reducer function to each element of the list.
// Returns the accumulated result.
func Reduce[TType any, TResult any](list *List[TType], reducer func(TResult, TType) TResult, initialValue TResult) TResult {
	var result = initialValue
	list.ForEach(func(index int, item TType) {
		result = reducer(result, item)
	})
	return result
}

// GroupBy groups the elements of the list by the specified key.
// Returns a map where the key is the result of the keySelector function
func GroupBy[TType any, TKey any](items *List[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *List[TType]] {
	var groups = hashmap.New[TKey, *List[TType]]()
	items.ForEach(func(index int, item TType) {
		var key = keySelector(item)
		if !groups.HasKey(key) {
			groups.Put(key, From(item))
		} else {
			groups.Get(key).Add(item)
		}
	})
	return groups
}
