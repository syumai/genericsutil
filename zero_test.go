package genericsutil

import "testing"

func TestIsZero(t *testing.T) {
	tests := map[string]struct {
		value int // comparable type value
		want  bool
	}{
		"zero": {
			value: 0,
			want:  true,
		},
		"not zero": {
			value: 1,
			want:  false,
		},
	}
	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := IsZero(tc.value)
			if tc.want != got {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}
