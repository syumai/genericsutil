package genericsutil

import (
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	i := 1
	wantInt := &i
	gotInt := PointerOf(1)
	if !reflect.DeepEqual(wantInt, gotInt) {
		t.Errorf("want true, got false")
	}

	str := "a"
	wantStr := &str
	gotStr := PointerOf("a")
	if !reflect.DeepEqual(wantStr, gotStr) {
		t.Errorf("want true, got false")
	}

	type X string
	x := X("a")
	wantX := &x
	gotX := PointerOf(X("a"))
	if !reflect.DeepEqual(wantX, gotX) {
		t.Errorf("want true, got false")
	}
}
