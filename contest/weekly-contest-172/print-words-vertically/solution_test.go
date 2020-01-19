package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input string
	want  []string
}

var cases = []Case{
	{input: "HOW ARE YOU", want: []string{"HAY", "ORO", "WEU"}},
	{input: "TO BE OR NOT TO BE", want: []string{"TBONTB", "OEROOE", "   T"}},
	{input: "CONTEST IS COMING", want: []string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"}},
}

func TestPrintVertically(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := printVertically(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
