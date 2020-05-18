package permutationinstring

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	type in struct {
		s1, s2 string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{s1: "ab", s2: "eidbaooo"}, want: true},
		{in: in{s1: "ab", s2: "eidboaoo"}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := checkInclusion(tt.in.s1, tt.in.s2)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
