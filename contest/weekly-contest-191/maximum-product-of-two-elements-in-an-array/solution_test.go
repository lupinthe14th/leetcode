package maximumproductoftwoelementsinanarray

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{3, 4, 5, 2}, want: 12},
		{in: []int{1, 5, 4, 5}, want: 16},
		{in: []int{3, 7}, want: 12},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxProduct(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
