package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input input
	want  []int
}

type input struct {
	restaurants                          [][]int
	veganFriendly, maxPrice, maxDistance int
}

var cases = []Case{
	{input: input{restaurants: [][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, veganFriendly: 1, maxPrice: 50, maxDistance: 10}, want: []int{3, 1, 5}},
	{input: input{restaurants: [][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, veganFriendly: 0, maxPrice: 50, maxDistance: 10}, want: []int{4, 3, 2, 1, 5}},
}

func TestFilterRestaurants(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := filterRestaurants(tt.input.restaurants, tt.input.veganFriendly, tt.input.maxPrice, tt.input.maxDistance)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
