package doctor

import "generic-collections/utils"

// RecoverDefaultFalse Recover from panic, return the default value of the type and false.
func RecoverDefaultFalse[T any]() (T, bool) {
	if r := recover(); r != nil {
		return utils.DefaultValue[T](), false
	}

	return utils.DefaultValue[T](), false
}

// RecoverFalse Recover from panic, return false.
func RecoverFalse() bool {
	if r := recover(); r != nil {
		return false
	}

	return false
}
