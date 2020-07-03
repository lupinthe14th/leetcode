package uniquepaths

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	type in struct {
		m, n int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{m: 3, n: 2}, want: 3},
		{in: in{m: 7, n: 3}, want: 28},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := uniquePaths(tt.in.m, tt.in.n)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
