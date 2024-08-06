package interfaces

type ISlicer[TType any] interface {
	Slice(index int, length int) IIndexableCollection[int, TType]
}
