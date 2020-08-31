package largestcomponentsizebycommonfactor

import (
	"fmt"
	"testing"
)

func TestLargestComponentSize(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{4, 6, 15, 35}, want: 4},
		{in: []int{20, 50, 9, 63}, want: 2},
		{in: []int{2, 3, 6, 7, 4, 12, 21, 39}, want: 8},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := largestComponentSize(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
