package urlshort

import (
	"testing"
)

func TestHelpMe(t *testing.T) {
	got := HelpMe()
	want := "me"

	if got != want {
		t.Errorf("Error: got %s, want %s", got, want)
	}
}
