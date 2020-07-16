package mypow

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {
	t.Parallel()
	type in struct {
		x float64
		n int
	}
	tests := []struct {
		in   in
		want float64
	}{
		{in: in{x: 2, n: 10}, want: 1024.0},
		{in: in{x: 2.1, n: 3}, want: 9.261000000000001},
		{in: in{x: 2, n: -2}, want: 0.25},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := myPow(tt.in.x, tt.in.n)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
