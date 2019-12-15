package sequentialdigits

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
	{id: 1, input: input{low: 100, high: 300}, want: []int{123, 234}},
	{id: 2, input: input{low: 1000, high: 13000}, want: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
}

func TestSequentialDigits(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := sequentialDigits(tt.input.low, tt.input.high)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
