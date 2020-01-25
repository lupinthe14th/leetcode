package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input uint32
	want  int
}

var cases = []Case{
	{input: 00000000000000000000000000001011, want: 3},
	{input: 00000000000000000000000010000000, want: 1},
}

func TestHammingWeight(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hammingWeight(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
