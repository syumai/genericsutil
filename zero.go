package genericsutil

// Zero returns zero value of given type argument type.
func Zero[T any]() (_ T) {
	return
}

// ZeroOf returns zero value of given type argument type.
// The type argument can be infered using function argument (function argument type inference).
func ZeroOf[T any](v T) (_ T) {
	return
}

