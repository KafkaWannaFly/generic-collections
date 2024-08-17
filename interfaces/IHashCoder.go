package interfaces

type IHashCoder interface {
	// HashCode returns the hash code of the object.
	// Two objects that are equal should have the same hash code.
	HashCode() string
}
