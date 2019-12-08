package climbingstairs

import (
	"fmt"
	"math/big"
	"testing"
)

var ansS1610 = "2136010781928948523844645283815574730385064008311995665747822755414246012661413902243436621593742528786413698527404458901613575702930324526113218162028403832301251481141629314362342496588380089335310516958008652225841325082508336642452960414578069346373381958171043406948199684066855095094569350148579342051690860548547807426339114290114"
var ans1610, _ = new(big.Int).SetString(ansS1610, 10)

var cases = []struct {
	id    int
	input int
	want  *big.Int
}{
	{id: 1, input: -1, want: big.NewInt(0)},
	{id: 2, input: 0, want: big.NewInt(1)},
	{id: 3, input: 1, want: big.NewInt(1)},
	{id: 4, input: 2, want: big.NewInt(2)},
	{id: 5, input: 3, want: big.NewInt(3)},
	{id: 6, input: 4, want: big.NewInt(5)},
	{id: 7, input: 5, want: big.NewInt(8)},
	{id: 8, input: 44, want: big.NewInt(1134903170)},
	{id: 9, input: 45, want: big.NewInt(1836311903)},
	{id: 10, input: 46, want: big.NewInt(2971215073)},
	{id: 11, input: 1610, want: ans1610},
}

func TestClimbStairs(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := climbStairs(tt.input)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
