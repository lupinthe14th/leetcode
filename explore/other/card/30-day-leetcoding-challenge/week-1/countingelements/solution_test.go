package countingelements

import (
	"fmt"
	"testing"
)

func TestCountElements(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 2, 3}, want: 2},
		{in: []int{1, 1, 3, 3, 5, 5, 7, 7}, want: 0},
		{in: []int{1, 3, 2, 3, 5, 0}, want: 3},
		{in: []int{1, 1, 2, 2}, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countElements(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
