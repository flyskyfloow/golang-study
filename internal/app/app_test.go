package app

import "testing"

func TestMessage(t *testing.T) {
	if got, want := Message(), "golang-study app is running"; got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}
