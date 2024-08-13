package guard

import (
	"fmt"
)

// EnsureIndexRange Provide a list with length and index, ensure the index is in the range of the list.
// Panic if index out of range.
//
// index: The index to check
//
// length: The length of the list
func EnsureIndexRange(index int, length int) {
	if index < 0 || index >= length {
		panic(fmt.Sprintf("Index %d is out of range for list of length %d", index, length))
	}
}
