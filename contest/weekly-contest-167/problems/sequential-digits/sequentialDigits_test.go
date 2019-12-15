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
	{id: 3, input: input{low: 10, high: 1000000000}, want: []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789, 1234, 2345, 3456, 4567, 5678, 6789, 12345, 23456, 34567, 45678, 56789, 123456, 234567, 345678, 456789, 1234567, 2345678, 3456789, 12345678, 23456789, 123456789}},
	{id: 4, input: input{low: 58, high: 155}, want: []int{67, 78, 89, 123}},
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
