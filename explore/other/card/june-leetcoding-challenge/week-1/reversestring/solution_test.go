package reversestring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		in, want []byte
	}{
		{in: []byte{'h', 'e', 'l', 'l', 'o'}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{in: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			reverseString(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Fatalf("in: %v want: %v", tt.in, tt.want)
			}
		})
	}
}
