# Methods

## `ICollection[TItem]` implementation

### `func (receiver *LinkedList[T]) Add(item T) interfaces.ICollection[T]`

`Add` adds an element to the linked list.

### `func (receiver *LinkedList[T]) AddAll(items interfaces.ICollection[T])`

`AddAll` adds all elements of the given collection to the linked list.

### `func (receiver *LinkedList[T]) Clear() interfaces.ICollection[T]`

`Clear` removes the specified item from the linked list. Returns the linked list itself.

### `func (receiver *LinkedList[T]) Clone() interfaces.ICollection[T]`

`Clone` returns a new linked list with the same elements.

### `func (receiver *LinkedList[T]) Count() int`

`Count` returns the number of elements in the linked list.

### `func (receiver *LinkedList[T]) Default() interfaces.ICollection[T]`

`Default` returns a default empty linked list.

{% hint style="info" %}
This method is used internally by the library. You might not need it.
{% endhint %}

### `func (receiver *LinkedList[T]) Filter(predicateFunc func(T) bool) interfaces.ICollection[T]`

`Filter` removes all elements from the linked list that do not satisfy the predicate function. Returns the linked list itself. Original linked list is not modified.

### `func (receiver *LinkedList[T]) ForEach(appliedFunc func(int, T))`

`ForEach` iterates over the elements of the linked list. The first argument of the `appliedFunc` is index of item. The second argument of the `appliedFunc` is the element of the linked list.

### `func (receiver *LinkedList[T]) Has(item T) bool`

`Has` checks if the linked list contains the specified item.

### `func (receiver *LinkedList[T]) HasAll(items interfaces.ICollection[T]) bool`

`HasAll` checks if the linked list contains all the items of the specified collection.

### `func (receiver *LinkedList[T]) HasAny(items interfaces.ICollection[T]) bool`

`HasAny` checks if the linked list contains any of the items of the specified collection.

### `func (receiver *LinkedList[T]) IsEmpty() bool`

`IsEmpty` checks if the linked list is empty.

### `func (receiver *LinkedList[T]) ToSlice() []T`

`ToSlice` converts the linked list to a slice.

## `IIndexableCollection[int, T]` Implementation

### `func (receiver *LinkedList[T]) GetAt(i int) T`

`GetAt` returns the value of the element at the specified index. Panics if the index is out of range.

### `func (receiver *LinkedList[T]) SetAt(i int, item T)`

`SetAt` update the value of the element at the specified index. Panics if the index is out of range.

### `func (receiver *LinkedList[T]) TryGetAt(i int) (T, bool)`

`TryGetAt` returns the value of the element at the specified index. Returns the value and true if the index is in range, otherwise the default value and false.

### `func (receiver *LinkedList[T]) TrySetAt(i int, item T) bool`

`TrySetAt` update the value of the element at the specified index. Returns true if the index is in range, otherwise false.

### `func (receiver *LinkedList[T]) AddFirst(item T) interfaces.ICollection[T]`

`AddFirst` adds the item to the beginning of the linked list. Returns the linked list itself.

### `func (receiver *LinkedList[T]) AddLast(item T) interfaces.ICollection[T]`

`AddLast` adds the item to the end of the linked list. Returns the linked list itself.

### `func (receiver *LinkedList[T]) AddBefore(i int, item T) interfaces.ICollection[T]`

`AddBefore` adds the item before the element at the specified index. Returns the linked list itself.

### `func (receiver *LinkedList[T]) TryAddBefore(i int, item T) bool`

`TryAddBefore` adds the item before the element at the specified index. Returns true if the index is in range, otherwise false.

### `func (receiver *LinkedList[T]) AddAfter(i int, item T) interfaces.ICollection[T]`

`AddAfter` adds the item after the element at the specified index. Returns the linked list itself.

### `func (receiver *LinkedList[T]) TryAddAfter(i int, item T) bool`

`TryAddAfter` adds the item after the element at the specified index. Returns true if the index is in range, otherwise false.

### `func (receiver *LinkedList[T]) RemoveAt(i int) T`

`RemoveAt` item at the specified index. Panics if the index is out of range.

### `func (receiver *LinkedList[T]) TryRemoveAt(i int) (T, bool)`

`TryRemoveAt` item at the specified index. Returns the value and true if the index is in range, otherwise the default value and false.

### `func (receiver *LinkedList[T]) RemoveFirst() T`

`RemoveFirst` item from the beginning of the linked list. Panics if the linked list is empty.

### `func (receiver *LinkedList[T]) RemoveLast() T`

`RemoveLast` item from the end of the linked list. Panics if the linked list is empty.

### `func (receiver *LinkedList[T]) FindFirst(predicate func(int, T) bool) int`

`FindFirst` the first element that satisfies the predicate. Returns the index of the element if found, otherwise -1.

### `func (receiver *LinkedList[T]) FindLast(predicate func(int, T) bool) int`

`FindLast` the last element that satisfies the predicate. Returns the index of the element if found, otherwise -1.

### `func (receiver *LinkedList[T]) FindAll(predicate func(int, T) bool) []int`

`FindAll` items based on predicate. Return all matched indexes.

### `func (receiver *LinkedList[T]) Slice(index int, length int) interfaces.IIndexableCollection[int, T]`

`Slice` returns a new collection that contains a slice of the original collection. The slice is defined by the index and length parameters. If the index is greater than the length of the collection or the length is negative, a panic will occur. If the length is greater than the remaining elements in the collection, the slice will continue to wrap around to the beginning of the collection. For example, if the linked list has elements \[1, 2, 3, 4, 5]:

\- Slice(2, 2) returns \[3, 4]

\- Slice(3, 3) returns \[4, 5, 1]

## `LinkedList[T]` specific methods

### `func (receiver *LinkedList[T]) Map(mapper func(int, T) any) *LinkedList[any]`

`Map` applies the given mapper function to each element of the linked list. Returns a new `LinkedList` containing the results. For better type assertion, use `linkedlist.Map`

### `func (receiver *LinkedList[T]) Reduce(reducer func(any, T) any, initialValue any) any`

`Reduce` applies the given reducer function to each element of the linked list. Returns the accumulated result. For better type assertion, use `linkedlist.Reduce`

### `func (receiver *LinkedList[T]) GroupBy(keySelector func(T) any) *hashmap.HashMap[any, *LinkedList[T]]`

`GroupBy` groups the elements of the linked list by the specified key. Returns a map where the key is the result of the `keySelector` function. For better type assertion, use `linkedlist.GroupBy`

### `func (receiver *LinkedList[T]) NodeAt(index int) *Node[T]`

`NodeAt` get Node object at certain index.

### `func (receiver *LinkedList[T]) TryNodeAt(index int) (*Node[T], bool)`

`TryNodeAt` get Node object at certain index. Return the Node object and true if index in range, else nil and false.
