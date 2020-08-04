package poweroffour

import (
	"fmt"
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   int
		want bool
	}{
		{in: 16, want: true},
		{in: 5, want: false},
		{in: 0, want: false},
		{in: 1, want: true},
		{in: 8, want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := isPowerOfFour(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
