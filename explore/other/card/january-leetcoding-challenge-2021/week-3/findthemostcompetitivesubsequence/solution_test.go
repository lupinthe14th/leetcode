package findthemostcompetitivesubsequence

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMostCompetitive(t *testing.T) {
	t.Parallel()
	type in struct {
		nums []int
		k    int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{nums: []int{3, 5, 2, 6}, k: 2}, want: []int{2, 6}},
		{in: in{nums: []int{2, 4, 3, 3, 5, 4, 9, 6}, k: 4}, want: []int{2, 3, 3, 4}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := mostCompetitive(tt.in.nums, tt.in.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
