package applydiscounteverynorders

import (
	"fmt"
	"testing"
)

func TestCashier(t *testing.T) {
	cashier := Constructor(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})

	type in struct {
		product []int
		amount  []int
	}
	type Case struct {
		in   in
		want float64
	}

	var cases = []Case{
		{in: in{product: []int{1, 2}, amount: []int{1, 2}}, want: 500.0},
		{in: in{product: []int{3, 7}, amount: []int{10, 10}}, want: 4000.0},
		{in: in{product: []int{1, 2, 3, 4, 5, 6, 7}, amount: []int{1, 1, 1, 1, 1, 1, 1}}, want: 800.0},
		{in: in{product: []int{4}, amount: []int{10}}, want: 4000.0},
		{in: in{product: []int{7, 3}, amount: []int{10, 10}}, want: 4000.0},
		{in: in{product: []int{7, 5, 3, 1, 6, 4, 2}, amount: []int{10, 10, 10, 9, 9, 9, 7}}, want: 7350.0},
		{in: in{product: []int{2, 3, 5}, amount: []int{5, 3, 2}}, want: 2500.0},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := cashier.GetBill(c.in.product, c.in.amount)
			if got != c.want {
				t.Errorf("%f, want: %f", got, c.want)
			}
		})
	}
}
