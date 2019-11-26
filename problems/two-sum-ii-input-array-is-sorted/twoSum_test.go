package twosumiiinputarrayissorted

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	numbers []int
	target  int
}

var cases = []struct {
	id    int
	input input
	want  []int
}{
	{id: 1, input: input{numbers: []int{2, 7, 11, 15}, target: 9}, want: []int{1, 2}},
	{id: 2, input: input{numbers: []int{2, 3, 4}, target: 6}, want: []int{1, 3}},
	{id: 3, input: input{numbers: []int{0, 0, 3, 4}, target: 0}, want: []int{1, 2}},
	{id: 4, input: input{numbers: []int{-3, 3, 4, 90}, target: 0}, want: []int{1, 2}},
	{id: 5, input: input{numbers: []int{5, 25, 75}, target: 100}, want: []int{2, 3}},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := twoSum(tt.input.numbers, tt.input.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
