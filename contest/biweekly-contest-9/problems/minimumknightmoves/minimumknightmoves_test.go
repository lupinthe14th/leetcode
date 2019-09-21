package minknightmoves

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id             int
	inputX, inputY int
	want           int
}{
	{id: 1, inputX: 2, inputY: 1, want: 1},
	{id: 2, inputX: 5, inputY: 5, want: 4},
}

func TestMinKnightMoves(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := minKnightMoves(tt.inputX, tt.inputY)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
