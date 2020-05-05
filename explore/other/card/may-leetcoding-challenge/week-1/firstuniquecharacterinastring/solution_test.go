package firstuniquecharacterinastring

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{in: "leetcode", want: 0},
		{in: "loveleetcode", want: 2},
		{in: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", want: -1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := firstUniqChar(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
