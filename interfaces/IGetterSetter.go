package interfaces

type IGetterSetter[TIndex any, TValue any] interface {
	Get(TIndex) TValue
	Set(TIndex, TValue)
}