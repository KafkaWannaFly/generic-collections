package interfaces

type ILesser[TType any] interface {
	// Less checks if the current item is less than the given item.
	Less(TType) bool
}
