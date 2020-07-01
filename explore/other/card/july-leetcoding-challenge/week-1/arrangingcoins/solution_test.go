package arrangingcoins

import (
	"fmt"
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 0, want: 0},
		{in: 1, want: 1},
		{in: 2, want: 1},
		{in: 5, want: 2},
		{in: 8, want: 3},
		{in: 1804289383, want: 60070},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := arrangeCoins(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})

	}
}
