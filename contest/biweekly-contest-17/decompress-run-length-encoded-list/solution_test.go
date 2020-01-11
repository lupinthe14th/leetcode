package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input []int
	want  []int
}

var cases = []Case{
	{input: []int{1, 2, 3, 4}, want: []int{2, 4, 4, 4}},
	{input: []int{42, 39}, want: []int{39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39}},
}

func TestDecompressRLElist(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := decompressRLElist(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
