package gc

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/guard"
	"github.com/KafkaWannaFly/generic-collections/interfaces"
)

// Slice returns a new collection that contains a slice of the original collection.
// The slice is defined by the index and length parameters.
// If the index is greater than the length of the collection or the length is negative, a panic will occur.
// If the length is greater than the remaining elements in the collection, the slice will continue wrap around to the beginning of the collection.
// For example, if the list has elements [1, 2, 3, 4, 5]:
//
// - Slice(2, 2) returns [3, 4]
//
// - Slice(3, 3) returns [4, 5, 1]
func Slice[TType any](collection interfaces.IIndexableCollection[int, TType], index int, length int) interfaces.IIndexableCollection[int, TType] {
	guard.EnsureIndexRange(index, collection.Count())

	if length < 0 || length > collection.Count() {
		panic(fmt.Sprintf("Length %d is out of range for list of length %d", length, collection.Count()))
	}

	out := (collection.Default().(any)).(interfaces.IIndexableCollection[int, TType])

	for i := 0; i < length; i++ {
		slicePos := i + index
		if i+index >= collection.Count() {
			slicePos = i + index - collection.Count()
		}

		out.Add(collection.GetAt(slicePos))
	}

	return out
}

// Map Internal function used by the library. Consider to use list.Map or set.Map instead.
func Map[TType any, TResult any](
	in interfaces.ICollection[TType],
	out interfaces.ICollection[TResult],
	mapper func(int, TType) TResult) interfaces.ICollection[TResult] {

	in.ForEach(func(index int, item TType) {
		out.Add(mapper(index, item))
	})

	return out
}

// Reduce Internal function used by the library. Consider to use list.Reduce or set.Reduce instead.
func Reduce[TType any, TResult any](
	list interfaces.ICollection[TType],
	reducer func(TResult, TType) TResult,
	initialValue TResult) TResult {

	result := initialValue
	list.ForEach(func(index int, item TType) {
		result = reducer(result, item)
	})

	return result
}

// Using reflection to implement GroupBy
// This is not the best way to implement GroupBy because it slow. Around x200 times slower than the explicit implementation.
// https://stackoverflow.com/questions/69909502/performance-of-using-reflect-to-access-struct-field-as-a-string-variable-vs-ac
//func GroupBy[TType any, TValue any](items interfaces.ICollection[TType], outMap any, keySelector func(TType) TValue) any {
//	hmValue := reflect.ValueOf(outMap)
//	items.ForEach(func(index int, item TType) {
//		key := keySelector(item)
//		if !hmValue.MethodByName("HasKey").Call([]reflect.Value{reflect.ValueOf(key)})[0].Bool() {
//			hmValue.MethodByName("Set").Call([]reflect.Value{reflect.ValueOf(key), reflect.ValueOf(items.Default().Add(item))})
//		} else {
//			var res interface{} = hmValue.MethodByName("Get").Call([]reflect.Value{reflect.ValueOf(key)})[0]
//			res.(interfaces.ICollection[TType]).Add(item)
//		}
//
//	})
//
//	return outMap
//}
