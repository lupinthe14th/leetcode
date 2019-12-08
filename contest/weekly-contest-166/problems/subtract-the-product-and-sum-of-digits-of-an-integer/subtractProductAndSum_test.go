package subtracttheproductandsumofdigitsofaninteger

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  int
}{
	{id: 1, input: 234, want: 15},
	{id: 2, input: 4421, want: 21},
}

func TestSubtractProductAndSum(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := subtractProductAndSum(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
