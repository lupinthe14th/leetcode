package minimumnumberofstepstomaketwostringsanagram

import (
	"fmt"
	"testing"
)

type input struct {
	s, t string
}
type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{s: "bab", t: "aba"}, want: 1},
	{input: input{s: "leetcode", t: "practice"}, want: 5},
	{input: input{s: "anagram", t: "mangaar"}, want: 0},
	{input: input{s: "xxyyzz", t: "xxyyzz"}, want: 0},
	{input: input{s: "friend", t: "family"}, want: 4},
}

func TestMinSteps(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minSteps(tt.input.s, tt.input.t)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
