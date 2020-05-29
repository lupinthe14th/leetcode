package countingbits

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		in   int
		want []int
	}{
		{in: 2, want: []int{0, 1, 1}},
		{in: 5, want: []int{0, 1, 1, 2, 1, 2}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countBits(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
