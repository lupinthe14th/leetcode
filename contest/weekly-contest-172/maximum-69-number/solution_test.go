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
	{input: 9669, want: 9969},
	{input: 9996, want: 9999},
	{input: 9999, want: 9999},
}

func TestMinInsertions(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maximum69Number(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
