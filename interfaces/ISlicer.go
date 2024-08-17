package interfaces

type ISlicer[TType any] interface {
	// Slice returns a new collection with items from the given index to the end of the collection.
	Slice(index int, length int) IIndexableCollection[int, TType]
}
