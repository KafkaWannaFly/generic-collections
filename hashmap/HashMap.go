package hashmap

import (
	"generic-collections/utils"
)

// HashMap is a collection that stores key-value pairs.
// If using struct as key, the struct must implement IHashCoder interface.
type HashMap[K any, V any] struct {
	elements map[string]*Entry[K, V]
	count    int
}

// New creates a new empty hashmap.
func New[K any, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{elements: make(map[string]*Entry[K, V])}
}

// From creates a new hashmap from a slice of entries.
func From[K any, V any](entries ...*Entry[K, V]) *HashMap[K, V] {
	var hashMap = New[K, V]()
	for _, entry := range entries {
		hashMap.elements[entry.HashCode()] = entry
	}

	hashMap.count = len(entries)

	return hashMap
}

// Of creates a new HashMap from a built-in map
func Of[K comparable, V any](inputMap map[K]V) *HashMap[K, V] {
	var result = New[K, V]()
	for k, v := range inputMap {
		result.Put(k, v)
	}

	return result
}

// region HashMap specific methods

// ForEach iterates over the elements of the hashmap.
// First argument of the appliedFunc is always 0 because hashmaps do not have indexes.
// Second argument of the appliedFunc is the entry of the hashmap.
func (receiver *HashMap[K, V]) ForEach(appliedFunc func(*Entry[K, V])) {
	for _, element := range receiver.elements {
		appliedFunc(element)
	}
}

// Add new element to the hashmap.
// If the element already exists, it is overwritten.
// Returns the hashmap itself.
func (receiver *HashMap[K, V]) Add(item *Entry[K, V]) *HashMap[K, V] {
	if !receiver.Has(item) {
		receiver.count++
	}

	receiver.elements[item.HashCode()] = item

	return receiver
}

// AddAll adds all elements of the given collection to the hashmap.
// Overwrites the element if it already exists.
// Returns the hashmap itself.
func (receiver *HashMap[K, V]) AddAll(items ...*Entry[K, V]) *HashMap[K, V] {
	for _, item := range items {
		receiver.Add(item)
	}

	receiver.count = len(receiver.elements)

	return receiver
}

// Count returns the number of elements in the hashmap.
func (receiver *HashMap[K, V]) Count() int {
	return receiver.count
}

// Has compare key and value of the item with the elements of the hashmap.
func (receiver *HashMap[K, V]) Has(item *Entry[K, V]) bool {
	var key = item.HashCode()
	value, ok := receiver.elements[key]
	return ok && utils.IsEqual(value, item)
}

// HasAll checks if all keys and values exist in the hashmap.
func (receiver *HashMap[K, V]) HasAll(items ...*Entry[K, V]) bool {
	var hasAll = true
	for _, item := range items {
		if !receiver.Has(item) {
			hasAll = false
		}
	}

	return hasAll
}

// HasAny checks if any key and value exist in the hashmap.
func (receiver *HashMap[K, V]) HasAny(items ...*Entry[K, V]) bool {
	var hasAny = false

	for _, item := range items {
		if receiver.Has(item) {
			hasAny = true
		}
	}

	return hasAny
}

// Clear removes all elements from the hashmap.
// Returns original hashmap itself.
func (receiver *HashMap[K, V]) Clear() *HashMap[K, V] {
	receiver.elements = make(map[string]*Entry[K, V])
	receiver.count = 0
	return receiver
}

// Filter removes the elements that do not satisfy the predicate.
// Return a new hashmap with the filtered elements.
// The original hashmap is not modified.
func (receiver *HashMap[K, V]) Filter(predicate func(*Entry[K, V]) bool) *HashMap[K, V] {
	var filtered = New[K, V]()
	receiver.ForEach(func(entry *Entry[K, V]) {
		if predicate(entry) {
			filtered.Add(entry)
		}
	})

	return filtered
}

// ToSlice converts the hashmap to a slice of entries.
func (receiver *HashMap[K, V]) ToSlice() []*Entry[K, V] {
	var slice = make([]*Entry[K, V], 0, receiver.Count())
	receiver.ForEach(func(entry *Entry[K, V]) {
		slice = append(slice, entry)
	})
	return slice
}

// IsEmpty checks if the hashmap is empty.
func (receiver *HashMap[K, V]) IsEmpty() bool {
	return receiver.Count() == 0
}

// Clone creates a new hashmap with the same elements.
func (receiver *HashMap[K, V]) Clone() *HashMap[K, V] {
	var cloned = New[K, V]()
	return cloned.AddAll(receiver.ToSlice()...)
}

// Put adds a new element to the hashmap. Similar to Add method.
// Returns the hashmap itself.
func (receiver *HashMap[K, V]) Put(key K, value V) *HashMap[K, V] {
	receiver.Add(&Entry[K, V]{
		Key:   key,
		Value: value,
	})

	return receiver
}

// GetKeys returns all keys of the hashmap.
func (receiver *HashMap[K, V]) GetKeys() []K {
	var keys = make([]K, 0, receiver.Count())
	receiver.ForEach(func(entry *Entry[K, V]) {
		keys = append(keys, entry.Key)
	})
	return keys
}

// GetValues returns all values of the hashmap.
func (receiver *HashMap[K, V]) GetValues() []V {
	var values = make([]V, 0, receiver.Count())
	receiver.ForEach(func(entry *Entry[K, V]) {
		values = append(values, entry.Value)
	})
	return values
}

// GetEntries returns all entries of the hashmap.
// Equivalent to ToSlice method.
func (receiver *HashMap[K, V]) GetEntries() []*Entry[K, V] {
	return receiver.ToSlice()
}

// HasKey checks if the key exists in the hashmap.
func (receiver *HashMap[K, V]) HasKey(key K) bool {
	var hashCode = utils.HashCodeOf(key)
	_, ok := receiver.elements[hashCode]
	return ok
}

// HasAllKey checks if all keys exist in the hashmap.
func (receiver *HashMap[K, V]) HasAllKey(keys []K) bool {
	var hasAll = true
	for _, key := range keys {
		if !receiver.HasKey(key) {
			hasAll = false
		}
	}

	return hasAll
}

// HasAnyKey checks if any key exists in the hashmap.
func (receiver *HashMap[K, V]) HasAnyKey(keys []K) bool {
	var hasAny = false
	for _, key := range keys {
		if receiver.HasKey(key) {
			hasAny = true
		}
	}

	return hasAny
}

// endregion

// region IIndexable implementation

// Get the value of the element at the specified key.
// If the key does not exist, default value of the value type is returned.
func (receiver *HashMap[K, V]) Get(index K) V {
	var key = utils.HashCodeOf(index)
	entry, ok := receiver.elements[key]

	if !ok {
		return utils.DefaultValue[V]()
	}

	return entry.Value
}

// Set the value of the element at the specified key.
// If the key does not exist, a new element is created with the specified key and value.
// If the key exists, the value of the element is updated.
func (receiver *HashMap[K, V]) Set(key K, value V) {
	receiver.Add(&Entry[K, V]{Key: key, Value: value})
}

// Find the key of the first element that satisfies the predicate.
func (receiver *HashMap[K, V]) Find(predicate func(entry *Entry[K, V]) bool) K {
	var key K
	receiver.ForEach(func(entry *Entry[K, V]) {
		if predicate(entry) {
			key = entry.Key
		}
	})

	return key
}

// Remove the element with the specified key.
// Returns the value of the removed element.
// If the key does not exist, the default value of the value type is returned.
func (receiver *HashMap[K, V]) Remove(key K) V {
	var hashCode = utils.HashCodeOf(key)
	var entry, ok = receiver.elements[hashCode]
	if ok {
		val := entry.Value

		delete(receiver.elements, hashCode)
		receiver.count--

		return val
	}

	return utils.DefaultValue[V]()
}

// endregion

// region Package functions

// IsHashMap checks if the collection is a hashmap.
func IsHashMap[K any, V any](collection any) bool {
	if collection == nil {
		return false
	}

	_, ok := collection.(*HashMap[K, V])
	return ok
}

// endregion
