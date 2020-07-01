package wordsearchii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	type in struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		in   in
		want []string
	}{
		{in: in{board: [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, words: []string{"oath", "pea", "eat", "rain"}}, want: []string{"eat", "oath"}},
	}
	for i, tt := range tests {
		i := i
		tt := tt
		t.Parallel()
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findWords(tt.in.board, tt.in.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
