package linkedlist

import (
	"github.com/KafkaWannaFly/generic-collections/gc"
	"github.com/KafkaWannaFly/generic-collections/hashmap"
)

// Map applies the given mapper function to each element of the list.
// Returns a new list containing the results. Don't modify the original list.
func Map[TType any, TResult any](linkedList *LinkedList[TType], mapper func(int, TType) TResult) *LinkedList[TResult] {
	return gc.Map(linkedList, New[TResult](), mapper).(*LinkedList[TResult])
}

// Reduce applies the given reducer function to each element of the list.
// Returns the accumulated result.
func Reduce[TType any, TResult any](linkedList *LinkedList[TType], reducer func(TResult, TType) TResult, initialValue TResult) TResult {
	return gc.Reduce(linkedList, reducer, initialValue)
}

// GroupBy groups the elements of the list by the specified key.
// Returns a map where the key is the result of the keySelector function
func GroupBy[TType any, TKey any](linkedList *LinkedList[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *LinkedList[TType]] {
	var groups = hashmap.New[TKey, *LinkedList[TType]]()
	linkedList.ForEach(func(index int, item TType) {
		var key = keySelector(item)
		if !groups.HasKey(key) {
			groups.Put(key, From[TType](item))
		} else {
			groups.Get(key).Add(item)
		}

	})
	return groups
}
