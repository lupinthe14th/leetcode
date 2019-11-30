package removeinterval

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	intervals   [][]int
	toBeRemoved []int
}

var cases = []struct {
	id    int
	input input
	want  [][]int
}{
	{id: 1, input: input{intervals: [][]int{{0, 2}, {3, 4}, {5, 7}}, toBeRemoved: []int{1, 6}}, want: [][]int{{0, 1}, {6, 7}}},
	{id: 2, input: input{intervals: [][]int{{0, 5}}, toBeRemoved: []int{2, 3}}, want: [][]int{{0, 2}, {3, 5}}},
}

func TestRemoveInterval(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := removeInterval(tt.input.intervals, tt.input.toBeRemoved)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
