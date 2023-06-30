package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello test", func(t *testing.T) {
		got := Hello("testy")
		want := "Hello, testy"

		if got != want {
			assertMessage(t, got, want)
		}
	})

	t.Run("empty string test", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			assertMessage(t, got, want)
		}
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
