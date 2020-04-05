package constructkpalindromestrings

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	type in struct {
		s string
		k int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{s: "annabelle", k: 2}, want: true},
		{in: in{s: "leetcode", k: 3}, want: false},
		{in: in{s: "true", k: 4}, want: true},
		{in: in{s: "yzyzyzyzyzyzyzy", k: 2}, want: true},
		{in: in{s: "cr", k: 7}, want: false},
		{in: in{s: "messi", k: 3}, want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canConstruct(tt.in.s, tt.in.k)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
