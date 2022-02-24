package genericsutil_test

import "github.com/syumai/genericsutil"

func ExampleAssertComparable() {
	// S1 implements comparable
	type S1 struct {
		A int
		B string
	}

	// S2 does not implement comparable
	// - []string does not implement comparable.
	type S2 struct {
		A int
		B []string
	}

	// OK
	type _ genericsutil.AssertComparable[S1]
	// NG => build error: S2 does not implement comparable
	// type _ genericsutil.AssertComparable[S2]

	// Output:
}
