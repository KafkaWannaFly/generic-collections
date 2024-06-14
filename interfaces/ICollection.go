package interfaces

type ICollection[T any] interface {
	IIterable[T]
	Add(T) ICollection[T]
	AddAll(ICollection[T]) ICollection[T]
	Count() int
	Contains(T) bool
	ContainsAll(ICollection[T]) bool
	Clear() ICollection[T]
	Filter(func(T) bool) ICollection[T]
	Get(any) T
	Set(any, T) ICollection[T]
	ToSlice() []T
	IsEmpty() bool
	Remove(T) ICollection[T]
	RemoveAll(ICollection[T]) ICollection[T]
}
