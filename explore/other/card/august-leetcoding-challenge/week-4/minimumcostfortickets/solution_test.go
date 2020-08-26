package minimumcostfortickets

import (
	"fmt"
	"testing"
)

func TestMincostTickets(t *testing.T) {
	t.Parallel()
	type in struct {
		days, costs []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{days: []int{1, 4, 6, 7, 8, 20}, costs: []int{2, 7, 15}}, want: 11},
		{in: in{days: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, costs: []int{2, 7, 15}}, want: 17},
		{in: in{days: []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 24, 25, 27, 28, 29, 30, 31, 34, 37, 38, 39, 41, 43, 44, 45, 47, 48, 49, 54, 57, 60, 62, 63, 66, 69, 70, 72, 74, 76, 78, 80, 81, 82, 83, 84, 85, 88, 89, 91, 93, 94, 97, 99}, costs: []int{9, 38, 134}}, want: 423},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := mincostTickets(tt.in.days, tt.in.costs)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
