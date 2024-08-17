package interfaces

// IIndexableCollection is an interface for a collection that has index
type IIndexableCollection[TIndex any, TValue any] interface {
	ICollection[TValue]
	IIndexableGetSet[TIndex, TValue]
	IIndexableAdder[TIndex, TValue]
	IIndexableRemover[TIndex, TValue]
	IIndexableFinder[TIndex, TValue]

	ISlicer[TValue]
}

// IIndexableGetSet is an interface for getting and setting item in a ICollection that have index
type IIndexableGetSet[TIndex any, TValue any] interface {
	// GetAt returns the item at the given index.
	GetAt(TIndex) TValue
	// SetAt sets the item at the given index.
	SetAt(TIndex, TValue)

	// TryGetAt returns the item at the given index and true if the index is valid; otherwise, it returns the default value and false.
	TryGetAt(TIndex) (TValue, bool)
	// TrySetAt sets the item at the given index and returns true if the index is valid; otherwise, it returns false.
	TrySetAt(TIndex, TValue) bool
}

// IIndexableAdder is an interface for adding items to a ICollection which has index
type IIndexableAdder[TIndex any, TValue any] interface {
	// AddFirst adds an item to the beginning of the collection.
	AddFirst(TValue) ICollection[TValue]
	// AddLast adds an item to the end of the collection.
	AddLast(TValue) ICollection[TValue]

	// AddBefore adds an item before the item at the given index.
	AddBefore(TIndex, TValue) ICollection[TValue]
	// TryAddBefore adds an item before the item at the given index and returns true if the index is valid; otherwise, it returns false.
	TryAddBefore(TIndex, TValue) bool

	// AddAfter adds an item after the item at the given index.
	AddAfter(TIndex, TValue) ICollection[TValue]
	// TryAddAfter adds an item after the item at the given index and returns true if the index is valid; otherwise, it returns false.
	TryAddAfter(TIndex, TValue) bool
}

// IIndexableRemover is an interface for removing items from a ICollection which has index
type IIndexableRemover[TIndex any, TValue any] interface {
	// RemoveFirst removes the first item from the collection.
	RemoveFirst() TValue
	// RemoveLast removes the last item from the collection.
	RemoveLast() TValue

	// RemoveAt removes the item at the given index.
	RemoveAt(TIndex) TValue
	// TryRemoveAt removes the item at the given index and returns true if the index is valid; otherwise, it returns the default value and false.
	TryRemoveAt(TIndex) (TValue, bool)
}

// IIndexableFinder is an interface for finding items in a ICollection which has index
type IIndexableFinder[TIndex any, TValue any] interface {
	// FindFirst returns the index of the first item that satisfies the given predicate.
	FindFirst(predicate func(TIndex, TValue) bool) TIndex
	// FindLast returns the index of the last item that satisfies the given predicate.
	FindLast(predicate func(TIndex, TValue) bool) TIndex
	// FindAll returns the indexes of all items that satisfy the given predicate.
	FindAll(predicate func(TIndex, TValue) bool) []TIndex
}
