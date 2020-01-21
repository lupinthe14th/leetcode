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
	{input: 2, want: []int{1, 1}},
	{input: 11, want: []int{2, 9}},
	{input: 10000, want: []int{1, 9999}},
	{input: 69, want: []int{1, 68}},
	{input: 1010, want: []int{11, 999}},
}

func TestGetNoZeroIntegers(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getNoZeroIntegers(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIsZeroInteger(t *testing.T) {
	var cases = []struct {
		input int
		want  bool
	}{
		{input: 0, want: true},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isZeroInteger(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
