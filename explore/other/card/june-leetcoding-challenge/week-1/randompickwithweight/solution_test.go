package randompickwithweight

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{in: []int{1}, want: []int{0}},
		// {in: []int{1, 3}, want: []int{0, 1, 1, 1, 0}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			s := Constructor(tt.in)

			for j := 0; j < len(tt.want); j++ {
				got := s.PickIndex()
				if got != tt.want[j] {
					t.Fatalf("got: %v want: %v", got, tt.want[j])
				}
			}
		})
	}
}
