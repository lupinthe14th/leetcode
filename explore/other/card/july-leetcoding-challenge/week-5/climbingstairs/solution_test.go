package climbingstairs

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want int
	}{
		{in: 2, want: 2},
		{in: 3, want: 3},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := climbStairs(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
