package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("A", 4)
	expected := "AAAA"
	if got != expected {
		t.Errorf("expected '%s', got '%s'", expected, got)
	}
}
