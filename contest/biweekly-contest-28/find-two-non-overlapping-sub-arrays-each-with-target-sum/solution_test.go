package findtwononoverlappingsubarrayseachwithtargetsum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinSumOfLengths(t *testing.T) {
	type in struct {
		arr    []int
		target int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{arr: []int{3, 2, 2, 4, 3}, target: 3}, want: 2},
		{in: in{arr: []int{7, 3, 4, 7}, target: 7}, want: 2},
		{in: in{arr: []int{4, 3, 2, 6, 2, 3, 4}, target: 6}, want: -1},
		{in: in{arr: []int{5, 5, 4, 4, 5}, target: 3}, want: -1},
		{in: in{arr: []int{3, 1, 1, 1, 5, 1, 2, 1}, target: 3}, want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minSumOfLengths(tt.in.arr, tt.in.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got :%v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
