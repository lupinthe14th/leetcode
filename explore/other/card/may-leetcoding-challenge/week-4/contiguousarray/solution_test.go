package contiguousarray

import (
	"fmt"
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{0, 1}, want: 2},
		{in: []int{0, 1, 0}, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findMaxLength(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
