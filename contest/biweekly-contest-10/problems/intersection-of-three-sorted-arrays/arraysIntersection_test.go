package intersectionofthreesortedarrays

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	arr1, arr2, arr3 []int
}

var cases = []struct {
	id    int
	input input
	want  []int
}{
	{id: 1, input: input{arr1: []int{1, 2, 3, 4, 5}, arr2: []int{1, 2, 5, 7, 9}, arr3: []int{1, 3, 4, 5, 8}}, want: []int{1, 5}},
}

func TestArraysIntersection(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := arraysIntersection(tt.input.arr1, tt.input.arr2, tt.input.arr3)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
