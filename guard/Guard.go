package guard

import (
	"fmt"
	"generic-collections/utils"
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

// RecoverDefaultFalse Recover from panic, return the default value of the type and false.
func RecoverDefaultFalse[T any]() (T, bool) {
	if r := recover(); r != nil {
		return utils.DefaultValue[T](), false
	}

	return utils.DefaultValue[T](), false
}

// RecoverFalse Recover from panic, return false.
func RecoverFalse[T any]() bool {
	if r := recover(); r != nil {
		return false
	}

	return false
}
