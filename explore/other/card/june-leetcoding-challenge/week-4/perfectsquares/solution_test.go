package perfectsquares

import (
	"fmt"
	"testing"
)

func TestNumSQuares(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 12, want: 3},
		{in: 13, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numSquares(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
