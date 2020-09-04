package repeatedsubstringpattern

import (
	"fmt"
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want bool
	}{
		{in: "abab", want: true},
		{in: "aba", want: false},
		{in: "abcabcabcabc", want: true},
		{in: "ababab", want: true},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := repeatedSubstringPattern(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
