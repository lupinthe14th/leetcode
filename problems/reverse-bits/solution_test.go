package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input, want uint32
}

var cases = []Case{
	{input: 43261596, want: 964176192},
	{input: 4294967293, want: 3221225471},
}

func TestReverseBits(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reverseBits(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
