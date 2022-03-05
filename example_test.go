package genericsutil_test

import (
	"fmt"

	"github.com/syumai/genericsutil"
)

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

func ExamplePointerOf() {
	pInt := genericsutil.PointerOf(100)
	fmt.Printf("value: %v, type: %T\n", *pInt, pInt)

	pStr := genericsutil.PointerOf("abcde")
	fmt.Printf("value: %v, type: %T\n", *pStr, pStr)

	type X string
	pX := genericsutil.PointerOf(X("abcde"))
	fmt.Printf("value: %v", *pX) // type is *main.X

	// Output:
	// value: 100, type: *int
	// value: abcde, type: *string
	// value: abcde
}
