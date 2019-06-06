package hey

import "testing"

func TestHi(t *testing.T) {
	want := "Hello, world."
	if got := Hi(); got != want {
		t.Errorf("Hi() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
