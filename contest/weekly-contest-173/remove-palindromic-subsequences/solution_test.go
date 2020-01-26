package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input string
	want  int
}

var cases = []Case{
	{input: "ababa", want: 1},
	{input: "abb", want: 2},
	{input: "baabb", want: 2},
	{input: "", want: 0},
	{input: "bbaabaaa", want: 2},
}

func TestRemovePalindromeSub(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := removePalindromeSub(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
