package gc

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/guard"
	"github.com/KafkaWannaFly/generic-collections/interfaces"
)

// Slice returns a new collection that contains a slice of the original collection.
// The slice is defined by the index and length parameters.
// If the index is greater than the length of the collection or the length is negative, a panic will occur.
// If the length is greater than the length of the collection, the slice will try to get the remaining items and then wrap around to the beginning of the collection.
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
