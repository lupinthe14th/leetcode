package checkifawordoccursasaprefixofanywordinasentence

import (
	"fmt"
	"testing"
)

func TestIsPrefixOfWord(t *testing.T) {
	type in struct {
		sentence, searchWord string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{sentence: "i love eating burger", searchWord: "burg"}, want: 4},
		{in: in{sentence: "this problem is an easy problem", searchWord: "pro"}, want: 2},
		{in: in{sentence: "i am tired", searchWord: "you"}, want: -1},
		{in: in{sentence: "i use triple pillow", searchWord: "pill"}, want: 4},
		{in: in{sentence: "hello from the other side", searchWord: "they"}, want: -1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPrefixOfWord(tt.in.sentence, tt.in.searchWord)
			if got != tt.want {
				t.Fatalf("in: %v got: %v tt.want: %v", tt.in, got, tt.want)
			}
		})
	}
}
