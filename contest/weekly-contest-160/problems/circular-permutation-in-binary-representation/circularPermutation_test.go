package circularpermutationinbinaryrepresentation

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	n, start int
}

var cases = []struct {
	id    int
	input input
	want  []int
}{
	{id: 1, input: input{n: 2, start: 3}, want: []int{3, 2, 0, 1}},
	{id: 2, input: input{n: 3, start: 2}, want: []int{2, 3, 1, 0, 4, 5, 7, 6}},
}

func TestCircularPermutation(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := circularPermutation(tt.input.n, tt.input.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
