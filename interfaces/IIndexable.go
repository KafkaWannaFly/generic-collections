package interfaces

type IIndexable[TIndex any, TValue any] interface {
	Get(TIndex) TValue
	Set(TIndex, TValue)
	Find(predicate func(TValue) bool) TIndex
	Remove(TIndex) TValue
}
