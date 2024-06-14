package interfaces

type IIterable[T any] interface {
	ForEach(func(int, T))
}
