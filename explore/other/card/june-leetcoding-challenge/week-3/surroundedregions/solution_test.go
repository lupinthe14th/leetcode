package surroundedregions

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		in, want [][]byte
	}{
		{
			in:   [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			want: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{in: [][]byte{}, want: [][]byte{}},
		{
			in:   [][]byte{{'X', 'X'}},
			want: [][]byte{{'X', 'X'}},
		},
		{
			in:   [][]byte{{'X', 'X', 'X'}, {'X', 'O', 'O'}, {'X', 'X', 'O'}},
			want: [][]byte{{'X', 'X', 'X'}, {'X', 'O', 'O'}, {'X', 'X', 'O'}},
		},
		{
			in:   [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}},
			want: [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}},
		},
		{
			in:   [][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}},
			want: [][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			solve(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Fatalf("got: %v want: %v", tt.in, tt.want)
			}
		})
	}
}
