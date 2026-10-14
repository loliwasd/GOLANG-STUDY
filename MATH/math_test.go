package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d, want %d", got, want)
	}

	got = Add(0, 0)
	want = 0
	if got != want {
		t.Errorf("Add(0, 0) = %d, want %d", got, want)
	}

	got = Add(-1, 1)
	want = 0
	if got != want {
		t.Errorf("Add(-1, 1) = %d, want %d", got, want)
	}

	got = Add(-2, -3)
	want = -5
	if got != want {
		t.Errorf("Add(-2, -3) = %d, want %d", got, want)
	}
}

func TestSub(t *testing.T) {

	got := Sub(-1, 1)
	want := -2
	if got != want {
		t.Errorf("Sub(-1, 1) = %d, want %d", got, want)
	}

	got = Sub(1, 100)
	want = -99
	if got != want {
		t.Errorf("Sub(1, 100) = %d, want %d", got, want)
	}

	got = Sub(5, 3)
	want = 2
	if got != want {
		t.Errorf("Sub(5, 3) = %d, want %d", got, want)
	}

}

func TestMax(t *testing.T) {

	got := Max(2, 3)
	want := 3
	if got != want {
		t.Errorf("Max(2,3) = %d, want %d", got, want)
	}

	got = Max(-2, -3)
	want = -2
	if got != want {
		t.Errorf("Max(-2,-3) = %d, want %d", got, want)
	}

	got = Max(0, 0)
	want = 0
	if got != want {
		t.Errorf("Max(0, 0) = %d, want %d", got, want)
	}
}
