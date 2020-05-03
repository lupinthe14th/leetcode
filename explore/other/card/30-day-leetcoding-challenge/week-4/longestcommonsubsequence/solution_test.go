package longestcommonsubsequence

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	type in struct {
		t1, t2 string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{t1: "abcde", t2: "ace"}, want: 3},
		{in: in{t1: "abc", t2: "abc"}, want: 3},
		{in: in{t1: "abc", t2: "def"}, want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := longestCommonSubsequence(tt.in.t1, tt.in.t2)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
