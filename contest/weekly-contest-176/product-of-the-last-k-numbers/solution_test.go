package productofthelastknumbers

import (
	"fmt"
	"testing"
)

func TestProductOfNumbers(t *testing.T) {
	productOfNumbers := Constructor()

	type in struct {
		action string
		num, k int
	}
	type Case struct {
		in   in
		want int
	}

	var cases = []Case{
		{in: in{action: "add", num: 3}},
		{in: in{action: "add", num: 0}},
		{in: in{action: "add", num: 2}},
		{in: in{action: "add", num: 5}},
		{in: in{action: "add", num: 4}},
		{in: in{action: "get", k: 2}, want: 20},
		{in: in{action: "get", k: 3}, want: 40},
		{in: in{action: "get", k: 4}, want: 0},
		{in: in{action: "add", num: 8}},
		{in: in{action: "get", k: 2}, want: 32},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			switch c.in.action {
			case "add":
				productOfNumbers.Add(c.in.num)
			case "get":
				got := productOfNumbers.GetProduct(c.in.k)
				if got != c.want {
					t.Errorf("%d, want: %d", got, c.want)
				}
			}
		})
	}
}
