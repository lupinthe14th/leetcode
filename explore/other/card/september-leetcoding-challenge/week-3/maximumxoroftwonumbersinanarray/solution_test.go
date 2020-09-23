package maximumxoroftwonumbersinanarray

import (
	"fmt"
	"testing"
)

func TestFindMaximumXOR(t *testing.T) {
	t.Parallel()
	type in struct {
		nums []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{nums: []int{3, 10, 5, 25, 2, 8}}, want: 28},
		{in: in{nums: []int{0}}, want: 0},
		{in: in{nums: []int{1}}, want: 0},
		{in: in{nums: []int{2, 4}}, want: 6},
		{in: in{nums: []int{8, 10, 2}}, want: 10},
		{in: in{nums: []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}}, want: 127},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := findMaximumXOR(tt.in.nums)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
