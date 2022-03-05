package genericsutil

// PointerOf returns pointer type value of given argument.
func PointerOf[T any](v T) *T {
	return &v
}

