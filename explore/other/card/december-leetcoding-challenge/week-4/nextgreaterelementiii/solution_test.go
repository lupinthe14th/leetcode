package nextgreaterelementiii

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want int
	}{
		{in: 12, want: 21},
		{in: 21, want: -1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := nextGreaterElement(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
