package strings

import (
	"testing"
)

func TestNewIterator(t *testing.T) {
	it := NewIterator("hello world")
	if got, want := it.content, "hello world"; got != want {
		t.Errorf("got %v want %v", got, want)
	}

	if got, want := it.pos, -1; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRead(t *testing.T) {
	it := NewIterator("hello")
	if got, want := it.Read(), "h"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := it.Read(), "e"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := it.Read(), "l"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := it.Read(), "l"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := it.Read(), "o"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
}

func TestDone(t *testing.T) {
	it := NewIterator("he")
	if _, got, want := it.Read(), it.Done(), false; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	if _, got, want := it.Read(), it.Done(), true; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
