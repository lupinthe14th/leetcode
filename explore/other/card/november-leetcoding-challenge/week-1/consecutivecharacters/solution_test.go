package consecutivecharacters

import (
	"fmt"
	"testing"
)

func TestMaxPower(t *testing.T) {
	t.Parallel()
	type in struct {
		s string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{s: "leetcode"}, want: 2},
		{in: in{s: "abbcccddddeeeeedcba"}, want: 5},
		{in: in{s: "triplepillooooow"}, want: 5},
		{in: in{s: "hooraaaaaaaaaaay"}, want: 11},
		{in: in{s: "tourist"}, want: 1},
		{in: in{s: "ccbccbb"}, want: 2},
		{in: in{s: "j"}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := maxPower(tt.in.s)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
