package subarraysumequalsk

import (
	"fmt"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	type in struct {
		nums []int
		k    int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{nums: []int{1, 1, 1}, k: 2}, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := subarraySum(tt.in.nums, tt.in.k)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}

		})
	}
}
