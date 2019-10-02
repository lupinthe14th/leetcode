package implementstrstr

import (
	"fmt"
	"testing"
)

type input struct {
	haystack string
	needle   string
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{haystack: "hello", needle: "ll"}, want: 2},
	{id: 2, input: input{haystack: "aaaaa", needle: "bba"}, want: -1},
	{id: 3, input: input{haystack: "bbbbb", needle: ""}, want: 0},
	{id: 4, input: input{haystack: "aaa", needle: "aaaa"}, want: -1},
	{id: 5, input: input{haystack: "aaa", needle: "aaa"}, want: 0},
}

func TestStrStr(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := strStr(tt.input.haystack, tt.input.needle)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
