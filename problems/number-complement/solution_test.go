package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input int
	want  int
}

var cases = []Case{
	{input: 5, want: 2},
	{input: 1, want: 0},
}

func TestFindComplement(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findComplement(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
