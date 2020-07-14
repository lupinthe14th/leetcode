package numberofsubstringswithonly1s

import (
	"fmt"
	"testing"
)

func TestNumSub(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{in: "0110111", want: 9},
		{in: "101", want: 2},
		{in: "111111", want: 21},
		{in: "000", want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			// t.Parallel()
			got := numSub(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
