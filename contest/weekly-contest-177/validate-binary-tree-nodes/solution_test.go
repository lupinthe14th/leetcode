package validatebinarytreenodes

import (
	"fmt"
	"testing"
)

type Case struct {
	in   in
	want bool
}

type in struct {
	n                     int
	leftChild, rightChild []int
}

var cases = []Case{
	{in: in{n: 4, leftChild: []int{1, -1, 3, -1}, rightChild: []int{2, -1, -1, -1}}, want: true},
	{in: in{n: 4, leftChild: []int{1, -1, 3, -1}, rightChild: []int{2, 3, -1, -1}}, want: false},
	{in: in{n: 2, leftChild: []int{1, 0}, rightChild: []int{-1, -1}}, want: false},
	{in: in{n: 6, leftChild: []int{1, -1, -1, 4, -1, -1}, rightChild: []int{2, -1, -1, 5, -1, -1}}, want: false},
}

func TestValidateBinaryTreeNodes(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := validateBinaryTreeNodes(c.in.n, c.in.leftChild, c.in.rightChild)
			if got != c.want {
				t.Errorf("%t, want: %t", got, c.want)
			}
		})
	}
}
