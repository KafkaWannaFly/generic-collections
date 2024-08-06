package interfaces

type ICollection[TItem any] interface {
	IIterable[TItem]
	Add(TItem) ICollection[TItem]
	AddAll(ICollection[TItem]) ICollection[TItem]
	Count() int
	Has(TItem) bool
	HasAll(ICollection[TItem]) bool
	HasAny(ICollection[TItem]) bool
	Clear() ICollection[TItem]
	Filter(func(TItem) bool) ICollection[TItem]
	ToSlice() []TItem
	IsEmpty() bool
	Clone() ICollection[TItem]
	Default() ICollection[TItem]
}
