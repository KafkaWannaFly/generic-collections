package hashmap

import (
	"generic-collections/interfaces"
	"generic-collections/utils"
)

type HashMap[K any, V any] struct {
	elements map[string]Entry[K, V]
	count    int
}

func New[K any, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{elements: make(map[string]Entry[K, V])}
}

func From[K any, V any](entries ...Entry[K, V]) *HashMap[K, V] {
	var hashMap = New[K, V]()
	for _, entry := range entries {
		hashMap.elements[entry.GetHashCode()] = entry
	}

	hashMap.count = len(entries)

	return hashMap
}

// region ICollection implementation

// ForEach iterates over the elements of the hashmap.
// First argument of the appliedFunc is always 0 because hashmaps do not have indexes.
// Second argument of the appliedFunc is the entry of the hashmap.
func (receiver *HashMap[K, V]) ForEach(appliedFunc func(int, Entry[K, V])) {
	for _, element := range receiver.elements {
		appliedFunc(0, element)
	}
}

func (receiver *HashMap[K, V]) Add(item Entry[K, V]) interfaces.ICollection[Entry[K, V]] {
	if !receiver.Has(item) {
		receiver.count++
	}

	receiver.elements[item.GetHashCode()] = item

	return receiver
}

// AddAll adds all elements of the given collection to the hashmap.
// Overwrites the element if it already exists.
func (receiver *HashMap[K, V]) AddAll(items interfaces.ICollection[Entry[K, V]]) interfaces.ICollection[Entry[K, V]] {
	items.ForEach(func(_ int, entry Entry[K, V]) {
		receiver.Add(entry)
	})

	receiver.count = len(receiver.elements)

	return receiver
}

func (receiver *HashMap[K, V]) Count() int {
	return receiver.count
}

// Has compare key and value of the item with the elements of the hashmap.
func (receiver *HashMap[K, V]) Has(item Entry[K, V]) bool {
	var key = item.GetHashCode()
	value, ok := receiver.elements[key]
	return ok && utils.IsEqual(value, item)
}

// HasAll checks if all keys and values exist in the hashmap.
func (receiver *HashMap[K, V]) HasAll(items interfaces.ICollection[Entry[K, V]]) bool {
	var hasAll = true
	items.ForEach(func(_ int, entry Entry[K, V]) {
		if !receiver.Has(entry) {
			hasAll = false
		}
	})

	return hasAll
}

func (receiver *HashMap[K, V]) Clear() interfaces.ICollection[Entry[K, V]] {
	receiver.elements = make(map[string]Entry[K, V])
	receiver.count = 0
	return receiver
}

// Filter removes the elements that do not satisfy the predicate.
// Return a new hashmap with the filtered elements.
// The original hashmap is not modified.
func (receiver *HashMap[K, V]) Filter(predicate func(Entry[K, V]) bool) interfaces.ICollection[Entry[K, V]] {
	var filtered = New[K, V]()
	receiver.ForEach(func(_ int, entry Entry[K, V]) {
		if predicate(entry) {
			filtered.Add(entry)
		}
	})

	return filtered
}

func (receiver *HashMap[K, V]) ToSlice() []Entry[K, V] {
	var slice = make([]Entry[K, V], 0, receiver.Count())
	receiver.ForEach(func(_ int, entry Entry[K, V]) {
		slice = append(slice, entry)
	})
	return slice
}

func (receiver *HashMap[K, V]) IsEmpty() bool {
	return receiver.Count() == 0
}

func (receiver *HashMap[K, V]) Clone() interfaces.ICollection[Entry[K, V]] {
	var cloned = New[K, V]()
	return cloned.AddAll(receiver)
}

// endregion

// region IIndexable implementation

// Get the value of the element at the specified key.
// If the key does not exist, default value of the value type is returned.
func (receiver *HashMap[K, V]) Get(index K) V {
	var key = utils.HashCodeOf(index)
	entry, _ := receiver.elements[key]

	return entry.Value
}

// Set the value of the element at the specified key.
// If the key does not exist, a new element is created with the specified key and value.
// If the key exists, the value of the element is updated.
func (receiver *HashMap[K, V]) Set(key K, value V) {
	receiver.Add(Entry[K, V]{Key: key, Value: value})
}

// Find the key of the first element that satisfies the predicate.
func (receiver *HashMap[K, V]) Find(predicate func(entry Entry[K, V]) bool) K {
	var key K
	receiver.ForEach(func(_ int, entry Entry[K, V]) {
		if predicate(entry) {
			key = entry.Key
		}
	})

	return key
}

// endregion

func (receiver *HashMap[K, V]) Put(key K, value V) interfaces.ICollection[Entry[K, V]] {
	receiver.Add(Entry[K, V]{
		Key:   key,
		Value: value,
	})

	return receiver
}

func (receiver *HashMap[K, V]) GetKeys() []K {
	var keys = make([]K, 0, receiver.Count())
	receiver.ForEach(func(_ int, entry Entry[K, V]) {
		keys = append(keys, entry.Key)
	})
	return keys
}

func (receiver *HashMap[K, V]) GetValues() []V {
	var values = make([]V, 0, receiver.Count())
	receiver.ForEach(func(_ int, entry Entry[K, V]) {
		values = append(values, entry.Value)
	})
	return values
}

func (receiver *HashMap[K, V]) GetEntries() []Entry[K, V] {
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

func (receiver *HashMap[K, V]) Remove(key K) V {
	var hashCode = utils.HashCodeOf(key)
	var entry, ok = receiver.elements[hashCode]
	if ok {
		delete(receiver.elements, hashCode)
		receiver.count--
		return entry.Value
	}

	return entry.Value
}
