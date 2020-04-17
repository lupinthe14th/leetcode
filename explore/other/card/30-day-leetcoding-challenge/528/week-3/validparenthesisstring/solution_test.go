package validparenthesisstring

import (
	"fmt"
	"testing"
)

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{in: "", want: true},
		{in: "()", want: true},
		{in: "()()()()()()()()()", want: true},
		{in: "()()()()(()()()()", want: false},
		{in: "((()", want: false},
		{in: "()))", want: false},
		{in: ")(", want: false},
		{in: "(()", want: false},
		{in: "*(", want: false},
		{in: ")*", want: false},
		{in: "((((((((()))))))))", want: true},
		{in: "(*)", want: true},
		{in: "*)", want: true},
		{in: "(*", want: true},
		{in: "*", want: true},
		{in: "(*))", want: true},
		{in: "((()))()(())(*()()())**(())()()()()((*()*))((*()*)", want: true},
		{in: "(())((())()()(*)(*()(())())())()()((()())((()))(*", want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := checkValidString(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
