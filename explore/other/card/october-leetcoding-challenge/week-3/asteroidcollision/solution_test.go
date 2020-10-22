package asteroidcollision

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	t.Parallel()
	type in struct {
		asteroids []int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{asteroids: []int{5, 10, -5}}, want: []int{5, 10}},
		{in: in{asteroids: []int{8, -8}}, want: []int{}},
		{in: in{asteroids: []int{10, 2, -5}}, want: []int{10}},
		{in: in{asteroids: []int{-2, -1, 1, 2}}, want: []int{-2, -1, 1, 2}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := asteroidCollision(tt.in.asteroids)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
