package sortcolors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		in, want []int
	}{
		{in: []int{2, 0, 2, 1, 1, 0}, want: []int{0, 0, 1, 1, 2, 2}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			sortColors(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Fatalf("got: %v want: %v", tt.in, tt.want)
			}
		})
	}
}
