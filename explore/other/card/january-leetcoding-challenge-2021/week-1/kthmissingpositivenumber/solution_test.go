package kthmissingpositivenumber

import (
	"fmt"
	"testing"
)

func TestFindKthPositive(t *testing.T) {
	t.Parallel()
	type in struct {
		arr []int
		k   int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{arr: []int{2, 3, 4, 7, 11}, k: 5}, want: 9},
		{in: in{arr: []int{1, 2, 3, 4}, k: 2}, want: 6},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := findKthPositive(tt.in.arr, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
