package intersectionofthreesortedarrays

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	low, high int
}

var cases = []struct {
	id    int
	input input
	want  []int
}{
	{id: 1, input: input{low: 0, high: 21}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 21}},
	{id: 2, input: input{low: 0, high: 101}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98, 101}},
}

func TestArraysIntersection(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := countSteppingNumbers(tt.input.low, tt.input.high)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
