package interfaces

type IIndexable[TIndex any, TValue any] interface {
	GetAt(TIndex) TValue
	SetAt(TIndex, TValue)
	Find(predicate func(TValue) bool) TIndex
	RemoveAt(TIndex) TValue

	TryGetAt(TIndex) (TValue, bool)
	TrySetAt(TIndex, TValue) bool
	TryRemoveAt(TIndex) (TValue, bool)
}
