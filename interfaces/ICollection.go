package interfaces

type ICollection[T any] interface {
	IIterable[T]
	Add(T) ICollection[T]
	AddAll(ICollection[T]) ICollection[T]
	Count() int
	Has(T) bool
	HasAll(ICollection[T]) bool
	HasAny(ICollection[T]) bool
	Clear() ICollection[T]
	Filter(func(T) bool) ICollection[T]
	ToSlice() []T
	IsEmpty() bool
	Clone() ICollection[T]
}
