# Methods

## `ICollection[TItem]` implementation

### `func (receiver *Set[T]) Add(item T) interfaces.ICollection[T]`

`Add` adds an element to the set. Overwrites the element if it already exists.

### `func (receiver *Set[T]) AddAll(items interfaces.ICollection[T])`

`AddAll` adds all elements of the given collection to the set. Overwrites the element if it already exists.

### `func (receiver *Set[T]) Clear() interfaces.ICollection[T]`

`Clear` removes the specified item from the set. Returns the set itself.

### `func (receiver *Set[T]) Clone() interfaces.ICollection[T]`

`Clone` returns a new set with the same elements.

### `func (receiver *Set[T]) Count() int`

`Count` returns the number of elements in the set.

### `func (receiver *Set[T]) Default() interfaces.ICollection[T]`

`Default` returns a default empty set.

{% hint style="info" %}
This method is used internally by the library. You might not need it.
{% endhint %}

### `func (receiver *Set[T]) Filter(predicateFunc func(T) bool) interfaces.ICollection[T]`

`Filter` removes all elements from the set that do not satisfy the predicate function. Returns the set itself. Original set is not modified.

### `func (receiver *Set[T]) ForEach(appliedFunc func(int, T))`

`ForEach` iterates over the elements of the set. The first argument of the `appliedFunc` is always 0 because sets do not have indexes. The second argument of the `appliedFunc` is the element of the set.

### `func (receiver *Set[T]) Has(item T) bool`

`Has` checks if the set contains the specified item.

### `func (receiver *Set[T]) HasAll(items interfaces.ICollection[T]) bool`

`HasAll` checks if the set contains all the items of the specified collection.

### `func (receiver *Set[T]) HasAny(items interfaces.ICollection[T]) bool`

`HasAny` checks if the set contains any of the items of the specified collection.

### `func (receiver *Set[T]) IsEmpty() bool`

`IsEmpty` checks if the set is empty.

### `func (receiver *Set[T]) ToSlice() []T`

`ToSlice` converts the set to a slice.
