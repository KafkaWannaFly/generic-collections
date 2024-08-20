# Methods

## `ICollection[TItem]` implementation

### `func (receiver *List[T]) Add(item T) interfaces.ICollection[T]`

`Add` adds an element to the list.

### `func (receiver *List[T]) AddAll(items interfaces.ICollection[T])`

`AddAll` adds all elements of the given collection to the list.

### `func (receiver *List[T]) Clear() interfaces.ICollection[T]`

`Clear` removes the specified item from the list. Returns the list itself.

### `func (receiver *List[T]) Clone() interfaces.ICollection[T]`

`Clone` returns a new list with the same elements.

### `func (receiver *List[T]) Count() int`

`Count` returns the number of elements in the list.

### `func (receiver *List[T]) Default() interfaces.ICollection[T]`

`Default` returns a default empty list.

{% hint style="info" %}
This method is used internally by the library. You might not need it.
{% endhint %}

### `func (receiver *List[T]) Filter(predicateFunc func(T) bool) interfaces.ICollection[T]`

`Filter` removes all elements from the list that do not satisfy the predicate function. Returns the list itself.
Original list is not modified.

### `func (receiver *List[T]) ForEach(appliedFunc func(int, T))`

`ForEach` iterates over the elements of the list. The first argument of the `appliedFunc` is index of item. The second
argument of the `appliedFunc` is the element of the list.

### `func (receiver *List[T]) Has(item T) bool`

`Has` checks if the list contains the specified item.

### `func (receiver *List[T]) HasAll(items interfaces.ICollection[T]) bool`

`HasAll` checks if the list contains all the items of the specified collection.

### `func (receiver *List[T]) HasAny(items interfaces.ICollection[T]) bool`

`HasAny` checks if the list contains any of the items of the specified collection.

### `func (receiver *List[T]) IsEmpty() bool`

`IsEmpty` checks if the list is empty.

### `func (receiver *List[T]) ToSlice() []T`

`ToSlice` converts the list to a slice.
