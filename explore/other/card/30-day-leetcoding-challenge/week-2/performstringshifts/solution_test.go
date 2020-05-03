package performstringshifts

import (
	"fmt"
	"testing"
)

func TestStringShift(t *testing.T) {
	type in struct {
		s     string
		shift [][]int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{s: "abc", shift: [][]int{{0, 1}, {1, 2}}}, want: "cab"},
		{in: in{s: "abcdefg", shift: [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}}, want: "efgabcd"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := stringShift(tt.in.s, tt.in.shift)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
