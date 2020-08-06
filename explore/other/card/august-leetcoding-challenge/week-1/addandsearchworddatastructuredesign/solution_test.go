package addandsearchworddatastructuredesign

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		act, word string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{act: "add", word: "bad"}},
		{in: in{act: "add", word: "dad"}},
		{in: in{act: "add", word: "mad"}},
		{in: in{act: "search", word: "pad"}, want: false},
		{in: in{act: "search", word: "bad"}, want: true},
		{in: in{act: "search", word: ".ad"}, want: true},
		{in: in{act: "search", word: "b.."}, want: true},
	}
	c := Constructor()
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			switch tt.in.act {
			case "add":
				c.AddWord(tt.in.word)
			case "search":
				got := c.Search(tt.in.word)
				if got != tt.want {
					t.Fatalf("got: %v want: %v", got, tt.want)
				}
			}
		})
	}
}
