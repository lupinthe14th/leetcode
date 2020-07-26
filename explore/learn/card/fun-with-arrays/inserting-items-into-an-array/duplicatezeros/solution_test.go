package duplicatezeros

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want []int
	}{
		{in: []int{1, 0, 2, 3, 0, 4, 5, 0}, want: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{in: []int{1, 2, 3}, want: []int{1, 2, 3}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			duplicateZeros(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Fatalf("in: %v want: %v", tt.in, tt.want)
			}
		})
	}
}
