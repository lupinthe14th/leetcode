package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input string
	want  string
}

var cases = []Case{
	{input: "10#11#12", want: "jkab"},
	{input: "1326#", want: "acz"},
	{input: "25#", want: "y"},
}

func TestFreqAlphabets(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := freqAlphabets(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
