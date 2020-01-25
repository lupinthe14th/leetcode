package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input int
	want  []int
}

var cases = []Case{
	{input: 0, want: []int{0}},
	{input: 1, want: []int{0, 1}},
	{input: 2, want: []int{0, 1, 1}},
	{input: 5, want: []int{0, 1, 1, 2, 1, 2}},
	{input: 8, want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
}

func TestCountBits(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countBits(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
