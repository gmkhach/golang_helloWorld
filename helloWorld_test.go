package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != "Hello, world." {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}