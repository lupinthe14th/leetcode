package wordsearch

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	t.Parallel()
	type in struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCCED"}, want: true},
		{in: in{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "SEE"}, want: true},
		{in: in{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCB"}, want: false},
		{in: in{board: [][]byte{{'a', 'b'}, {'c', 'd'}}, word: "abcd"}, want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := exist(tt.in.board, tt.in.word)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
