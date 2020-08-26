package sortarraybyparity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want []int
	}{
		{in: []int{3, 1, 2, 4}, want: []int{2, 4, 3, 1}},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortArrayByParity(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
