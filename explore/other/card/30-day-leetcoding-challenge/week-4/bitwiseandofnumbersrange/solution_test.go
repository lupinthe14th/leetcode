package bitwiseandofnumbersrange

import (
	"fmt"
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	type in struct {
		m, n int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{m: 5, n: 7}, want: 4},
		{in: in{m: 5, n: 8}, want: 0},
		{in: in{m: 0, n: 1}, want: 0},
		{in: in{m: 0, n: 2147483647}, want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := rangeBitwiseAnd(tt.in.m, tt.in.n)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
