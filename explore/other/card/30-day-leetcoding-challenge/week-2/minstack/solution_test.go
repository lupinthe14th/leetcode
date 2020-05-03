package minstack

import "testing"

func TestMinStack(t *testing.T) {
	minstack := Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	t.Run("1", func(t *testing.T) {
		got := minstack.GetMin()
		if got != -3 {
			t.Fatalf("got: %v, want: %v", got, -3)
		}
	})
	minstack.Pop()
	t.Run("2", func(t *testing.T) {
		got := minstack.Top()
		if got != 0 {
			t.Fatalf("got: %v, want: %v", got, 0)
		}
	})
	t.Run("3", func(t *testing.T) {
		got := minstack.GetMin()
		if got != -2 {
			t.Fatalf("got: %v, want: %v", got, -2)
		}
	})
}
