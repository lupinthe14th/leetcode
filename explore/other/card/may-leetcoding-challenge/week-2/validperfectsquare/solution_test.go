package validperfectsquare

import (
	"fmt"
	"testing"
)

func TestIsPerfextSquare(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{in: 16, want: true},
		{in: 14, want: false},
		{in: 100, want: true},
		{in: 0, want: true},
		{in: 1, want: true},
		{in: 2, want: false},
		{in: 3, want: false},
		{in: 4, want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPerfectSquare(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
