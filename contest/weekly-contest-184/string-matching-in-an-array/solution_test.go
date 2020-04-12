package stringmatchinginanarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringMatching(t *testing.T) {
	tests := []struct {
		in, want []string
	}{
		{in: []string{"mass", "as", "hero", "superhero"}, want: []string{"as", "hero"}},
		{in: []string{"leetcode", "et", "code"}, want: []string{"et", "code"}},
		{in: []string{"blue", "green", "bu"}, want: []string{}},
		{in: []string{"leetcoder", "leetcode", "od", "hamlet", "am"}, want: []string{"leetcode", "od", "am"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := stringMatching(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsSubstring(t *testing.T) {
	type in struct {
		s, ss string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{s: "mass", ss: "as"}, want: true},
		{in: in{s: "superhero", ss: "hero"}, want: true},
		{in: in{s: "as", ss: "mass"}, want: false},
		{in: in{s: "blue", ss: "bu"}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isSubstring(tt.in.s, tt.in.ss)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
