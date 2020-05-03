package maximumsubarray

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxSubArray(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
