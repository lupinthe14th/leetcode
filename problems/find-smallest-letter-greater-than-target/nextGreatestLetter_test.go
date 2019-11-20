package findsmallestlettergreaterthantarget

import (
	"fmt"
	"testing"
)

type input struct {
	letters []byte
	target  byte
}

var cases = []struct {
	id    int
	input input
	want  byte
}{
	{id: 1, input: input{letters: []byte{'c', 'f', 'j'}, target: 'a'}, want: 'c'},
	{id: 2, input: input{letters: []byte{'c', 'f', 'j'}, target: 'c'}, want: 'f'},
	{id: 3, input: input{letters: []byte{'c', 'f', 'j'}, target: 'd'}, want: 'f'},
	{id: 4, input: input{letters: []byte{'c', 'f', 'j'}, target: 'g'}, want: 'j'},
	{id: 5, input: input{letters: []byte{'c', 'f', 'j'}, target: 'j'}, want: 'c'},
	{id: 6, input: input{letters: []byte{'c', 'f', 'j'}, target: 'k'}, want: 'c'},
}

func TestNextGreatestLetter(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := nextGreatestLetter(tt.input.letters, tt.input.target)
			if got != tt.want {
				t.Errorf("%s, want: %s", string(got), string(tt.want))
			}
		})
	}
}
