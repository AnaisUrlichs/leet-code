package greetings

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 4)
	want := 5

	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}

}
