package hashmap

import (
	"generic-collections/utils"
)

type Entry[K any, V any] struct {
	Key   K
	Value V
}

func NewEntry[K any, V any](key K, value V) Entry[K, V] {
	return Entry[K, V]{Key: key, Value: value}
}

func (receiver Entry[K, V]) Clone() Entry[K, V] {
	return NewEntry(receiver.Key, receiver.Value)
}

func (receiver Entry[K, V]) GetHashCode() string {
	return utils.HashCodeOf(receiver.Key)
}

func (receiver Entry[K, V]) Equals(entry Entry[K, V]) bool {
	return receiver.GetHashCode() == entry.GetHashCode()
}
