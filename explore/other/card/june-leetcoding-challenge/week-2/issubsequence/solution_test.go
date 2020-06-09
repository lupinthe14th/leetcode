package issubsequence

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	type in struct {
		s, t string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{s: "abc", t: "ahbgdc"}, want: true},
		{in: in{s: "axc", t: "ahbgdc"}, want: false},
		{in: in{s: "abc", t: "bahbgdca"}, want: true},
		{in: in{s: "abc", t: "bahgdcb"}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isSubsequence(tt.in.s, tt.in.t)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
