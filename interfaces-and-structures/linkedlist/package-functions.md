# Package Functions

These functions live under `linkedlist` package.

### `func New[T any]() *LinkedList[T]`

`New` creates a new empty linkedlist.

```go
emptyIntegerLinkedList := linkedlist.New[int]()
```

### `func From[T any](items ...T) *LinkedList[T]`

`From` creates a new linkedlist from a slice of elements.

```go
integerLinkedList := linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
```

### `func IsLinkedList[T any](item any) bool`

`IsLinkedList` checks if the specified item is a linkedlist of type T.

```go
integerLinkedList := linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

linkedlist.IsLinkedList[int](integerLinkedList) // true
linkedlist.IsLinkedList[string](integerLinkedList) // false
```

###
`func Map[TType any, TResult any](linkedlist *LinkedList[TType], mapper func(int, TType) TResult) *LinkedList[TResult]`

`Map` applies the given mapper function to each element of the linkedlist. Returns a new linkedlist containing the
results.

###
`func Reduce[TType any, TResult any](linkedlist *LinkedList[TType], reducer func(TResult, TType) TResult, initialValue TResult) TResult`

`Reduce` applies the given reducer function to each element of the linkedlist. Returns the accumulated result.

###
`func GroupBy[TType any, TKey any](linkedlist *LinkedList[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *LinkedList[TType]]`

`GroupBy` groups the elements of the linkedlist by the specified key. Returns a map where the key is the result of the
`keySelector` function.
