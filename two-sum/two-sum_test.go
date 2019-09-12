package main

import (
	"fmt"
	"reflect"
	"testing"
)

type args struct {
	nums   []int
	target int
}

var cases = []struct {
	input args
	want  []int
}{
	{input: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
	{input: args{nums: []int{2, 7, 11, 15}, target: 26}, want: []int{2, 3}},
}

func TestTwoSumBruteForce(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			actual := twoSumBruteForce(tt.input.nums, tt.input.target)
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("%v, want %v", actual, tt.want)
			}
		})
	}
}
