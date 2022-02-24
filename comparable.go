package genericsutil

// AssertComparable asserts the given type argument's type implements comparable interface at compile time.
type AssertComparable[_ comparable] struct{}
