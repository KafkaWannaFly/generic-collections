package interfaces

// IIndexableGetSet is an interface for getting and setting item in a ICollection that have index
type IIndexableGetSet[TIndex any, TValue any] interface {
	GetAt(TIndex) TValue
	SetAt(TIndex, TValue)

	TryGetAt(TIndex) (TValue, bool)
	TrySetAt(TIndex, TValue) bool
}

// IIndexableAdder is an interface for adding items to a ICollection which has index
type IIndexableAdder[TIndex any, TValue any] interface {
	AddFirst(TValue) ICollection[TValue]
	AddLast(TValue) ICollection[TValue]

	AddBefore(TIndex, TValue) ICollection[TValue]
	TryAddBefore(TIndex, TValue) bool

	AddAfter(TIndex, TValue) ICollection[TValue]
	TryAddAfter(TIndex, TValue) bool
}

// IIndexableRemover is an interface for removing items from a ICollection which has index
type IIndexableRemover[TIndex any, TValue any] interface {
	RemoveFirst() TValue
	RemoveLast() TValue

	RemoveAt(TIndex) TValue
	TryRemoveAt(TIndex) (TValue, bool)
}

// IIndexableFinder is an interface for finding items in a ICollection which has index
type IIndexableFinder[TIndex any, TValue any] interface {
	FindFirst(predicate func(TValue) bool) TIndex
	FindLast(predicate func(TValue) bool) TIndex
	FindAll(predicate func(TValue) bool) []TIndex
}
