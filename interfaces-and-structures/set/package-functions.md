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

### `func IsSet[T any](collection any) bool`

`IsSet` checks if the specified item is a set of type T.

```go
integerSet := set.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

set.IsSet[int](integerSet) // true
set.IsSet[string](integerSet) // false
```
