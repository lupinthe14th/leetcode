package main

import (
	"fmt"
	"testing"
)

type input struct {
	a, b, c int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{a: 2, b: 6, c: 5}, want: 3},
	{input: input{a: 4, b: 2, c: 7}, want: 1},
	{input: input{a: 1, b: 2, c: 3}, want: 0},
	{input: input{a: 8, b: 3, c: 5}, want: 3},
	//	{input: input{a: 1e9, b: 1e9, c: 1e9}, want: 0},
}

func TestMinFlips(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minFlips(tt.input.a, tt.input.b, tt.input.c)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
