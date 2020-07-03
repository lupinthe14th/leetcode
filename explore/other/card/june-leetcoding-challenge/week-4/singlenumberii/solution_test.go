package singlenumberii

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1}, want: 1},
		{in: []int{2, 2, 3, 2}, want: 3},
		{in: []int{0, 1, 0, 1, 0, 1, 99}, want: 99},
		{in: []int{0, 1, 0, 99, 0, 99, 99}, want: 1},
		{in: []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}, want: -4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := singleNumber(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
