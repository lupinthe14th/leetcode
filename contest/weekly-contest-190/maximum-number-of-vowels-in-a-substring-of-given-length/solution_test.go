package maximumnumberofvowelsinasubstringofgivenlength

import (
	"fmt"
	"testing"
)

func TestMaxVowels(t *testing.T) {
	type in struct {
		s string
		k int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{s: "abciiidef", k: 3}, want: 3},
		{in: in{s: "aeiou", k: 2}, want: 2},
		{in: in{s: "leetcode", k: 3}, want: 2},
		{in: in{s: "rhythms", k: 4}, want: 0},
		{in: in{s: "tryhard", k: 4}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxVowels(tt.in.s, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})

	}
}
