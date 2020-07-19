package numberofnodesinthesubtreewiththesamelabel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountSubTrees(t *testing.T) {
	t.Parallel()
	type in struct {
		n      int
		edges  [][]int
		labels string
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, labels: "abaedcd"}, want: []int{2, 1, 1, 1, 1, 1, 1}},
		{in: in{n: 4, edges: [][]int{{0, 1}, {1, 2}, {0, 3}}, labels: "bbbb"}, want: []int{4, 2, 1, 1}},
		{in: in{n: 5, edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, labels: "aababb"}, want: []int{3, 2, 1, 1, 1}},
		{in: in{n: 6, edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {3, 4}, {4, 5}}, labels: "cbabaa"}, want: []int{1, 2, 1, 1, 2, 1}},
		{in: in{n: 7, edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}, labels: "aaabaaa"}, want: []int{6, 5, 4, 1, 3, 2, 1}},
		{in: in{n: 4, edges: [][]int{{0, 2}, {0, 3}, {1, 2}}, labels: "aeed"}, want: []int{1, 1, 2, 1}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			//	t.Parallel()
			got := countSubTrees(tt.in.n, tt.in.edges, tt.in.labels)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
