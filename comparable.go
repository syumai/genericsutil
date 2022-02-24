package genericsutil

// AssertComparable asserts the given type argument implements comparable interface at compile time.
type AssertComparable[_ comparable] struct{}
