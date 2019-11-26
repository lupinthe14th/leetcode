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
