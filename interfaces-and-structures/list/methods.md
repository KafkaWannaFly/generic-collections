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

`Filter` removes all elements from the list that do not satisfy the predicate function. Returns the list itself. Original list is not modified.

### `func (receiver *List[T]) ForEach(appliedFunc func(int, T))`

`ForEach` iterates over the elements of the list. The first argument of the `appliedFunc` is index of item. The second argument of the `appliedFunc` is the element of the list.

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



## `IIndexableCollection[int, T]` Implementation

### `func (receiver *List[T]) GetAt(i int) T`

`GetAt` returns the value of the element at the specified index. Panics if the index is out of range.

### `func (receiver *List[T]) SetAt(i int, item T)`

`SetAt` update the value of the element at the specified index. Panics if the index is out of range.

### `func (receiver *List[T]) TryGetAt(i int) (T, bool)`

`TryGetAt` returns the value of the element at the specified index. Returns the value and true if the index is in range, otherwise the default value and false.

### `func (receiver *List[T]) TrySetAt(i int, item T) bool`

`TrySetAt` update the value of the element at the specified index. Returns true if the index is in range, otherwise false.



### `func (receiver *List[T]) AddFirst(item T) interfaces.ICollection[T]`

`AddFirst` adds the item to the beginning of the list. Returns the list itself.

### `func (receiver *List[T]) AddLast(item T) interfaces.ICollection[T]`

`AddLast` adds the item to the end of the list. Returns the list itself.

### `func (receiver *List[T]) AddBefore(i int, item T) interfaces.ICollection[T]`

`AddBefore` adds the item before the element at the specified index. Returns the list itself.

### `func (receiver *List[T]) TryAddBefore(i int, item T) bool`

`TryAddBefore` adds the item before the element at the specified index. Returns true if the index is in range, otherwise false.

### `func (receiver *List[T]) AddAfter(i int, item T) interfaces.ICollection[T]`

`AddAfter` adds the item after the element at the specified index. Returns the list itself.

### `func (receiver *List[T]) TryAddAfter(i int, item T) bool`

`TryAddAfter` adds the item after the element at the specified index. Returns true if the index is in range, otherwise false.



### `func (receiver *List[T]) RemoveAt(i int) T`

`RemoveAt` item at the specified index. Panics if the index is out of range.

### `func (receiver *List[T]) TryRemoveAt(i int) (T, bool)`

`TryRemoveAt` item at the specified index. Returns the value and true if the index is in range, otherwise the default value and false.

### `func (receiver *List[T]) RemoveFirst() T`

`RemoveFirst` item from the beginning of the list. Panics if the list is empty.

### `func (receiver *List[T]) RemoveLast() T`

`RemoveLast` item from the end of the list. Panics if the list is empty.



### `func (receiver *List[T]) FindFirst(predicate func(int, T) bool) int`

`FindFirst` the first element that satisfies the predicate. Returns the index of the element if found, otherwise -1.

### `func (receiver *List[T]) FindLast(predicate func(int, T) bool) int`

`FindLast` the last element that satisfies the predicate. Returns the index of the element if found, otherwise -1.

### `func (receiver *List[T]) FindAll(predicate func(int, T) bool) []int`

`FindAll` items based on predicate. Return all matched indexes.



### `func (receiver *List[T]) Slice(index int, length int) interfaces.IIndexableCollection[int, T]`

`Slice` returns a new collection that contains a slice of the original collection. The slice is defined by the index and length parameters. If the index is greater than the length of the collection or the length is negative, a panic will occur. If the length is greater than the remaining elements in the collection, the slice will continue to wrap around to the beginning of the collection. For example, if the list has elements \[1, 2, 3, 4, 5]:&#x20;

\- Slice(2, 2) returns \[3, 4]&#x20;

\- Slice(3, 3) returns \[4, 5, 1]



## `List[T]` specific methods

### `func (receiver *List[T]) Map(mapper func(int, T) any) *List[any]`

`Map` applies the given mapper function to each element of the list. Returns a new List containing the results. For better type assertion, use `list.Map`

### `func (receiver *List[T]) Reduce(reducer func(any, T) any, initialValue any) any`

`Reduce` applies the given reducer function to each element of the list. Returns the accumulated result. For better type assertion, use `list.Reduce`

### `func (receiver *List[T]) GroupBy(keySelector func(T) any) *hashmap.HashMap[any, *List[T]]`

`GroupBy` groups the elements of the list by the specified key. Returns a map where the key is the result of the `keySelector` function. For better type assertion, use `list.GroupBy`
