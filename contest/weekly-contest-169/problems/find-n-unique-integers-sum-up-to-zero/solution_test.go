package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  []int
}{
	{id: 1, input: 5, want: []int{-2, -1, 0, 1, 2}},
	{id: 2, input: 3, want: []int{-1, 0, 1}},
	{id: 3, input: 1, want: []int{0}},
	{id: 4, input: 4, want: []int{-2, -1, 1, 2}},
}

func TestSumZero(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := sumZero(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
