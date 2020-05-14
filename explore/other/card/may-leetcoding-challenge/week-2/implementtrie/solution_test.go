package implementtrie

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		action, name string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{action: "ins", name: "apple"}},
		{in: in{action: "ser", name: "apple"}, want: true},
		{in: in{action: "ser", name: "app"}, want: false},
		{in: in{action: "sw", name: "app"}, want: true},
		{in: in{action: "ins", name: "app"}},
		{in: in{action: "ser", name: "app"}, want: true},
	}

	trie := Constructor()
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var got bool
			switch tt.in.action {
			case "ins":
				trie.Insert(tt.in.name)
			case "ser":
				got = trie.Search(tt.in.name)
			case "sw":
				got = trie.StartsWith(tt.in.name)

			}
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
