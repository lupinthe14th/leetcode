package findtheduplicatenumber

import (
	"fmt"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 3, 4, 2, 2}, want: 2},
		{in: []int{3, 1, 3, 4, 2}, want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findDuplicate(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
