package gc

import (
	"generic-collections/guard"
	"generic-collections/interfaces"
)

// Slice returns a new collection that contains a slice of the original collection.
// The slice is defined by the index and length parameters.
// If the length is positive, the slice will start at the index and end at the index + length. If index + length is greater than the collection length, the slice will end at the collection length.
// If the length is negative, the slice will start at the index - length and end at the index.
func Slice[TType any](collection interfaces.IIndexableCollection[int, TType], index int, length int) interfaces.IIndexableCollection[int, TType] {
	guard.EnsureIndexRange(index, collection.Count())

	out := (collection.Default().(any)).(interfaces.IIndexableCollection[int, TType])
	if length == 0 {
		return out
	} else if length > 0 {
		collection.ForEach(func(i int, item TType) {
			if i >= index && i < index+length {
				out.Add(item)
			}
		})
	} else {
		// TODO: Implement negative length
	}

	return out
}
