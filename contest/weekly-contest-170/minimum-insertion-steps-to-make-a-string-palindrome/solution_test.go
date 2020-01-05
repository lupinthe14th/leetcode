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
	{input: "zzazz", want: 0},
	{input: "mbadm", want: 2},
	{input: "leetcode", want: 5},
	{input: "g", want: 0},
	{input: "no", want: 1},
}

func TestMinInsertions(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minInsertions(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
