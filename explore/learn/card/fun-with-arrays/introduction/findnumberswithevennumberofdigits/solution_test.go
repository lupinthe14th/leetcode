package findnumberswithevennumberofdigits

import (
	"fmt"
	"testing"
)

func TestFindNumbers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{12, 345, 2, 6, 7896}, want: 2},
		{in: []int{555, 901, 482, 1771}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := findNumbers(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
