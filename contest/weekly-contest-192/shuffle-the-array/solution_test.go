package shufflethearray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	type in struct {
		nums []int
		n    int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{nums: []int{2, 5, 1, 3, 4, 7}, n: 3}, want: []int{2, 3, 5, 4, 1, 7}},
		{in: in{nums: []int{1, 2, 3, 4, 4, 3, 2, 1}, n: 4}, want: []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{in: in{nums: []int{1, 1, 2, 2}, n: 2}, want: []int{1, 2, 1, 2}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := shuffle(tt.in.nums, tt.in.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
