package uniquebinarysearchtrees

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 1, want: 1},
		{in: 2, want: 2},
		{in: 3, want: 5},
		{in: 4, want: 14},
		{in: 5, want: 42},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numTrees(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
