package interfaces

type ICollection[T any] interface {
	IIterable[T]
	Add(T)
	AddAll(ICollection[T])
	Count() int
	Contains(T) bool
	ContainsAll(ICollection[T]) bool
	Clear()
	Filter(func(T) bool) ICollection[T]
	Get(any) T
	Set(any, T)
	ToSlice() []T
	IsEmpty() bool
	Remove(T)
	RemoveAll(ICollection[T])
}
