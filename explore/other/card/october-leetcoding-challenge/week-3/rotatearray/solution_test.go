package rotatearray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	t.Parallel()
	type in struct {
		nums []int
		k    int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{in: in{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 10}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{in: in{nums: []int{-1, -100, 3, 99}, k: 2}, want: []int{3, 99, -1, -100}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			rotate(tt.in.nums, tt.in.k)
			if !reflect.DeepEqual(tt.in.nums, tt.want) {
				t.Fatalf("in: %v want: %v", tt.in, tt.want)
			}
		})
	}
}
