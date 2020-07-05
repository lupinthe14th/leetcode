package uglynumberii

import (
	"fmt"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 10, want: 12},
		{in: 11, want: 15},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := nthUglyNumber(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
