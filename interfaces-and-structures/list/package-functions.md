---
hidden: true
---

# Package Functions

These functions live under `list` package.

### `func New[T any]() *List[T]`

`New` creates a new empty list.

```go
emptyIntegerList := list.New[int]()
```

### `func From[T any](items ...T) *List[T]`

`From` creates a new list from a slice of elements.

```go
integerList := list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
```

### `func IsList[T any](item any) bool`

`IsList` checks if the specified item is a list of type T.

```go
integerList := list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

list.IsList[int](integerList) // true
list.IsList[string](integerList) // false
```

### `func Map[TType any, TResult any](list *List[TType], mapper func(int, TType) TResult) *List[TResult]`

`Map` applies the given mapper function to each element of the list. Returns a new list containing the results.

### `func Reduce[TType any, TResult any](list *List[TType], reducer func(TResult, TType) TResult, initialValue TResult) TResult`

`Reduce` applies the given reducer function to each element of the list. Returns the accumulated result.

### `func GroupBy[TType any, TKey any](list *List[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *List[TType]]`

`GroupBy` groups the elements of the list by the specified key. Returns a map where the key is the result of the `keySelector` function.
