package main

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
	{id: 1, input: []int{17, 18, 5, 4, 6, 1}, want: []int{18, 6, 6, 6, 1, -1}},
}

func TestSolution(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := solution(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
