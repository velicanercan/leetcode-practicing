package validParentheses

import "testing"

func Test_validParentheses(t *testing.T) {
	t.Run("isValid", func(t *testing.T) {
		s := "()"
		got := isValid(s)
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("isValid", func(t *testing.T) {
		s := "()[]{}"
		got := isValid(s)
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("isValid", func(t *testing.T) {
		s := "([])"
		got := isValid(s)
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
