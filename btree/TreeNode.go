package btree

import "generic-collections/interfaces"

type TreeNode[T any] struct {
	interfaces.IComparer[TreeNode[T]]
	interfaces.IHashCoder

	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}
