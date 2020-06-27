package thekthfactorofn

import (
	"fmt"
	"testing"
)

func TestKthFactor(t *testing.T) {
	type in struct {
		n, k int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 12, k: 3}, want: 3},
		{in: in{n: 7, k: 2}, want: 7},
		{in: in{n: 4, k: 4}, want: -1},
		{in: in{n: 1, k: 1}, want: 1},
		{in: in{n: 1000, k: 3}, want: 4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := kthFactor(tt.in.n, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
