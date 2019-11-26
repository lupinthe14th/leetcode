package arraytransformation

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  []int
}{
	{id: 1, input: []int{6, 2, 3, 4}, want: []int{6, 3, 3, 4}},
	{id: 2, input: []int{1, 6, 3, 4, 3, 5}, want: []int{1, 4, 4, 4, 4, 5}},
	{id: 3, input: []int{2, 1, 2, 1, 1, 2, 2, 1}, want: []int{2, 2, 1, 1, 1, 2, 2, 1}},
}

func TestTransformArray(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := transformArray(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
