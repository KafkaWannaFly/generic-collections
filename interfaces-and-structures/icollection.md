# ICollection

`ICollection[TItem]` is a shared interface which is implemented by: [list](list/ "mention"), [set](set/ "mention"), [linkedlist.md](linkedlist.md "mention"), [stack.md](stack.md "mention"), [queue.md](queue.md "mention"), [binary-tree.md](binary-tree.md "mention"). Declaration:

```go
package interfaces

type ICollection[TItem any] interface {
	// ForEach iterates over the collection and applies the given function to each item.
	// The function receives the index of the item and the item itself.
	ForEach(func(int, TItem))
	
	// Add adds an item to the collection.
	Add(TItem) ICollection[TItem]
	
	// AddAll adds all items from the given collection to the collection.
	AddAll(ICollection[TItem]) ICollection[TItem]
	
	// Count returns the number of items in the collection.
	Count() int
	
	// Has checks if the collection contains the given item.
	Has(TItem) bool
	
	// HasAll checks if the collection contains all items from the given collection.
	HasAll(ICollection[TItem]) bool
	
	// HasAny checks if the collection contains any item from the given collection.
	HasAny(ICollection[TItem]) bool
	
	// Clear removes all items from the collection.
	Clear() ICollection[TItem]
	
	// Filter returns a new collection with items that satisfy the given function.
	// The original collection remains unchanged.
	Filter(func(TItem) bool) ICollection[TItem]
	
	// ToSlice returns a slice with all items from the collection.
	ToSlice() []TItem
	
	// IsEmpty checks if the collection is empty.
	IsEmpty() bool
	
	// Clone returns a new collection with the same items.
	Clone() ICollection[TItem]
	
	// Default returns a new empty collection of the same type.
	// This method is created for internal use only.
	Default() ICollection[TItem]
}
```
