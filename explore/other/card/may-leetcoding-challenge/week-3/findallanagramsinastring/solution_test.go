package findallanagramsinastring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	type in struct {
		s, p string
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{s: "cbaebabacd", p: "abc"}, want: []int{0, 6}},
		{in: in{s: "abab", p: "ab"}, want: []int{0, 1, 2}},
		{in: in{s: "af", p: "be"}, want: []int{}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findAnagrams(tt.in.s, tt.in.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
