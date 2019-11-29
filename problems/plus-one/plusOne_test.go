package plusone

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
	{id: 1, input: []int{1, 2, 3}, want: []int{1, 2, 4}},
	{id: 2, input: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
	{id: 3, input: []int{9}, want: []int{1, 0}},
	{id: 4, input: []int{9, 9, 9, 9, 9, 9, 9}, want: []int{1, 0, 0, 0, 0, 0, 0, 0}},
}

func TestPlusOne(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := plusOne(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
