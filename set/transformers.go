package set

import "github.com/KafkaWannaFly/generic-collections/hashmap"

// Map applies the given mapper function to each element of the list.
// Returns a new list containing the results.
func Map[TType any, TResult any](set *Set[TType], mapper func(int, TType) TResult) *Set[TResult] {
	var result = New[TResult]()
	set.ForEach(func(index int, item TType) {
		result.Add(mapper(index, item))
	})
	return result
}

// Reduce applies the given reducer function to each element of the list.
// Returns the accumulated result.
func Reduce[TType any, TResult any](set *Set[TType], reducer func(TResult, TType) TResult, initialValue TResult) TResult {
	var result = initialValue
	set.ForEach(func(index int, item TType) {
		result = reducer(result, item)
	})
	return result
}

// GroupBy groups the elements of the list by the specified key.
// Returns a map where the key is the result of the keySelector function
func GroupBy[TType any, TKey any](set *Set[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *Set[TType]] {
	var groups = hashmap.New[TKey, *Set[TType]]()
	set.ForEach(func(index int, item TType) {
		var key = keySelector(item)
		if !groups.HasKey(key) {
			groups.Put(key, From[TType](item))
		} else {
			groups.Get(key).Add(item)
		}

	})
	return groups
}
