package utils

import (
	"fmt"
	"generic-collections/interfaces"
)

// IsEqual If a, b implement IComparer, then use Compare method to compare them.
// If a, b implement IHashCoder, then use GetHashCode method to compare them.
// Else. use == operand
func IsEqual[T any](a T, b T) bool {
	var ia interface{} = a
	var ib interface{} = b

	iComparerA, aOk := ia.(interfaces.IComparer[T])

	if aOk {
		return iComparerA.Compare(b) == 0
	}

	iHashCoder, aOk := ia.(interfaces.IHashCoder)
	if aOk {
		return iHashCoder.GetHashCode() == HashCodeOf(b)
	}

	return ia == ib
}

// HashCodeOf If item implement IHashCoder, then use GetHashCode method to get hash code.
// Else. use fmt.Sprintf to convert item to string
func HashCodeOf[T any](item T) string {
	var iItem interface{} = item
	var iHashCoder, ok = iItem.(interfaces.IHashCoder)

	if ok {
		return iHashCoder.GetHashCode()
	}

	return fmt.Sprintf("%v", item)
}
