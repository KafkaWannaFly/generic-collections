package utils

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/interfaces"
)

// IsEqual If a, b implement IComparer, then use Compare method to compare them.
// If a, b implement IHashCoder, then use HashCode method to compare them.
// Else. use fmt.Sprintf to convert item to string
func IsEqual[T any](a T, b T) bool {
	return HashCodeOf(a) == HashCodeOf(b)
}

// HashCodeOf If item implement IHashCoder, then use HashCode method to get hash code.
// Else. use fmt.Sprintf to convert item to string
func HashCodeOf[T any](item T) string {
	var iItem interface{} = item
	var iHashCoder, ok = iItem.(interfaces.IHashCoder)

	if ok {
		return iHashCoder.HashCode()
	}

	return fmt.Sprintf("%v", item)
}

func DefaultValue[T any]() T {
	var result T
	return result
}
