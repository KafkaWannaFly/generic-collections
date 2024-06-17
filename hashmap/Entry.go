package hashmap

import "generic-collections/interfaces"

type Entry[K any, V any] struct {
	interfaces.IComparer[Entry[K, V]]
	interfaces.IHashCoder

	key   K
	value V
}
