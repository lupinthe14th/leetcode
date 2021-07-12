package bulbswitcheriv

import (
	"fmt"
	"testing"
)

func TestMinFlips(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want int
	}{
		{in: "10111", want: 3},
		{in: "101", want: 3},
		{in: "00000", want: 0},
		{in: "001011101", want: 5},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minFlips(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
