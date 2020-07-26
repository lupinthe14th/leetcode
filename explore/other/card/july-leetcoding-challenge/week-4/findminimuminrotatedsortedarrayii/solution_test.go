package findminimuminrotatedsortedarrayii

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 3, 5}, want: 1},
		{in: []int{2, 2, 2, 0, 1}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findMin(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
