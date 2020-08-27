package fizzbuzz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   int
		want []string
	}{
		{in: 15, want: []string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		},
		},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := fizzBuzz(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
