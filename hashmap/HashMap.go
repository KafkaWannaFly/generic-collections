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
	receiver.elements[item.GetHashCode()] = item
	receiver.count++
	return receiver
}

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

func (receiver *HashMap[K, V]) Has(item Entry[K, V]) bool {
	var key = item.GetHashCode()
	_, ok := receiver.elements[key]
	return ok
}

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

// region IGetterSetter implementation

func (receiver *HashMap[K, V]) Get(index K) V {
	var key = utils.HashCodeOf(index)
	entry, _ := receiver.elements[key]

	return entry.Value
}

// Set the value of the element at the specified key.
// If the key does not exist, a new element is created with the specified key and value.
// If the key exists, the value of the element is updated.
func (receiver *HashMap[K, V]) Set(index K, value V) {
	receiver.Add(Entry[K, V]{Key: index, Value: value})
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
