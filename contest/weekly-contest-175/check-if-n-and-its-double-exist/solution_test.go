package checkifnanditsdoubleexist

import (
	"fmt"
	"testing"
)

type Case struct {
	input []int
	want  bool
}

var cases = []Case{
	{input: []int{10, 2, 5, 3}, want: true},
	{input: []int{7, 1, 14, 11}, want: true},
	{input: []int{3, 1, 7, 11}, want: false},
	{input: []int{-2, 0, 10, -19, 4, 6, -8}, want: false},
}

func TestCheckIfExist(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := checkIfExist(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
