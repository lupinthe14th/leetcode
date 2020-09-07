package wordpattern

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	t.Parallel()
	type in struct {
		pattern, str string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{pattern: "abba", str: "dog cat cat dog"}, want: true},
		{in: in{pattern: "abba", str: "dog cat cat fish"}, want: false},
		{in: in{pattern: "aaaa", str: "dog cat cat dog"}, want: false},
		{in: in{pattern: "abba", str: "dog dog dog dog"}, want: false},
		{in: in{pattern: "aba", str: "cat cat cat dog"}, want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := wordPattern(tt.in.pattern, tt.in.str)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
