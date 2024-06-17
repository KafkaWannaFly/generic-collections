package hashmap

import "generic-collections/interfaces"

type HashMap[K any, V any] struct {
	interfaces.ICollection[Entry[K, V]]
	interfaces.IGetterSetter[string, Entry[K, V]]

	elements map[string]Entry[K, V]
	count    int
}
