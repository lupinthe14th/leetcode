package distributecandiestopeople

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	t.Parallel()
	type in struct {
		candies, num_people int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{candies: 7, num_people: 4}, want: []int{1, 2, 3, 1}},
		{in: in{candies: 10, num_people: 3}, want: []int{5, 2, 3}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := distributeCandies(tt.in.candies, tt.in.num_people)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want %v", tt.in, got, tt.want)
			}
		})
	}
}
