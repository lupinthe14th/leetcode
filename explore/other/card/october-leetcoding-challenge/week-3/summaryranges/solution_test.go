package summaryranges

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	t.Parallel()
	type in struct {
		nums []int
	}
	tests := []struct {
		in   in
		want []string
	}{
		{in: in{nums: []int{0, 1, 2, 4, 5, 7}}, want: []string{"0->2", "4->5", "7"}},
		{in: in{nums: []int{0, 2, 3, 4, 6, 8, 9}}, want: []string{"0", "2->4", "6", "8->9"}},
		{in: in{nums: []int{}}, want: []string{}},
		{in: in{nums: []int{-1}}, want: []string{"-1"}},
		{in: in{nums: []int{0}}, want: []string{"0"}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := summaryRanges(tt.in.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
