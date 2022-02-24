package genericsutil

// Zero returns zero value of given type argument type.
func Zero[T any]() (_ T) {
	return
}

// ZeroOf returns zero value of given type argument type.
// The type argument can be inferred using function argument (function argument type inference).
func ZeroOf[T any](v T) (_ T) {
	return
}

// IsZero compares given function argument to its type's zero value and returns result as a boolean value.
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

