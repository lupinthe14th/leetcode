package singleelementinasortedarray

import (
	"fmt"
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 1, 2, 3, 3, 4, 4, 8, 8}, want: 2},
		{in: []int{3, 3, 7, 7, 10, 11, 11}, want: 10},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := singleNonDuplicate(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
