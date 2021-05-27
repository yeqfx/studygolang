package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	want := 3.0
	got := Add(1, 2)
	if got != want {
		t.Errorf("Add(%#v) = %0.0f, want %0.0f", got, got, want)
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(8, 4) != 4 {
		t.Error("8 - 4 did not equal 4")
	}
}
