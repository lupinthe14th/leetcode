package hindexii

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {

	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{0, 1, 3, 5, 6}, want: 3},
		{in: []int{3, 4, 5, 8, 10}, want: 4},
		{in: []int{3, 3, 5, 8, 25}, want: 3},
		{in: []int{1}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hIndex(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
