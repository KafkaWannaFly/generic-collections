package utils

import (
	"generic-collections/interfaces"
)

// IsEqual If a, b implement IComparer, then use Compare method to compare them.
// Else. use == operand
func IsEqual[T any](a T, b T) bool {
	var ia interface{} = a
	var ib interface{} = b

	ica, aOk := ia.(interfaces.IComparer[T])

	if aOk {
		return ica.Compare(b) == 0
	}

	return ia == ib
}
