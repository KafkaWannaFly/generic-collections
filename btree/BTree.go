package btree

import "generic-collections/interfaces"

type BTree[T any] struct {
	interfaces.ICollection[T]

	root  *TreeNode[T]
	count int
}
