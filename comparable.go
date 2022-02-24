package genericsutil

// IsComparable checks if given type argument implements comparable interface in compile time.
type IsComparable[_ comparable] struct{}
