package prisoncellsafterndays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrisonAfterNDays(t *testing.T) {
	type in struct {
		cells []int
		n     int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{cells: []int{0, 1, 0, 1, 1, 0, 0, 1}, n: 7}, want: []int{0, 0, 1, 1, 0, 0, 0, 0}},
		{in: in{cells: []int{1, 0, 0, 1, 0, 0, 1, 0}, n: 100}, want: []int{0, 1, 0, 1, 0, 0, 1, 0}},
		{in: in{cells: []int{1, 0, 0, 1, 0, 0, 1, 0}, n: 1000000000}, want: []int{0, 0, 1, 1, 1, 1, 1, 0}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := prisonAfterNDays(tt.in.cells, tt.in.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
