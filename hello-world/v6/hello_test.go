package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chairs", "")
		want := "Hello,Chairs"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello,world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola,Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour,Lauren"
		assertCorrectMessage(t, got, want)
	})
}
