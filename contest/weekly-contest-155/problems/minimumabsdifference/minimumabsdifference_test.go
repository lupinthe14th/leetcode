package minimumabsdifference

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  [][]int
}{
	{id: 1, input: []int{4, 2, 1, 3}, want: [][]int{{1, 2}, {2, 3}, {3, 4}}},
	{id: 2, input: []int{1, 3, 6, 10, 15}, want: [][]int{{1, 3}}},
	{id: 3, input: []int{3, 8, -10, 23, 19, -4, -14, 27}, want: [][]int{{-14, -10}, {19, 23}, {23, 27}}},
}

func TestMinimumAbsDifference(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := minimumAbsDifference(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want; %v", got, tt.want)
			}
		})
	}

}
