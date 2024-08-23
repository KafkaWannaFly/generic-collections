# Package Functions

These functions live under `set` package.

### `func New[T any]() *Set[T]`

`New` creates a new empty set.

```go
emptyIntegerSet := set.New[int]()
```

### `func From[T any](items ...T) *Set[T]`

`From` creates a new set from a slice of elements.

```go
integerSet := set.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
```

### `func IsSet[T any](item any) bool`

`IsSet` checks if the specified item is a set of type T.

```go
integerSet := set.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

set.IsSet[int](integerSet) // true
set.IsSet[string](integerSet) // false
```

### `func Map[TType any, TResult any](set *Set[TType], mapper func(int, TType) TResult) *Set[TResult]`

`Map` applies the given mapper function to each element of the list. Returns a new list containing the results.

### `func Reduce[TType any, TResult any](set *Set[TType], reducer func(TResult, TType) TResult, initialValue TResult) TResult`

`Reduce` applies the given reducer function to each element of the list. Returns the accumulated result.

### `func GroupBy[TType any, TKey any](set *Set[TType], keySelector func(TType) TKey) *hashmap.HashMap[TKey, *Set[TType]]`

`GroupBy` groups the elements of the list by the specified key. Returns a map where the key is the result of the `keySelector` function.
