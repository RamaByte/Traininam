package mathx

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("Add(2, 3) = %d, want %d", got, want)
	}
}

func TestZero(t *testing.T) {
	got := Add(0, 0)
	want := 0

	if got != want {
		t.Errorf("Add(0, 0) = %d, want %d", got, want)
	}
}
