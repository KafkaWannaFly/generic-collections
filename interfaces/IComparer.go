package interfaces

type IComparer[T any] interface {
	// Compare with other element
	//
	// Return 0 if elements are equal,  1 if element is greater than others, -1 if element is lesser than others.
	Compare(T) int
}
