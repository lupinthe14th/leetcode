package longestsubarrayof1safterdeletingoneelement

import (
	"fmt"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 1, 0, 1}, want: 3},
		{in: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, want: 5},
		{in: []int{1, 1, 1}, want: 2},
		{in: []int{1, 1, 0, 0, 1, 1, 1, 0, 1}, want: 4},
		{in: []int{0, 0, 0}, want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := longestSubarray(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
