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

func TestDiv(t *testing.T) {
	testCases := []struct {
		name    string
		a       int
		b       int
		want    int
		wanterr bool
	}{
		{"1 by 1", 1, 1, 1, false},
		{"5 by 2", 5, 2, 2, false},
		{"div by 0", 1, 0, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Div(tc.a, tc.b)
			//if got, wanterr != tc.want, tc.wanterr {
			if tc.wanterr {
				if err == nil {
					t.Fatalf("Div(%d, %d): expected error, got nil", tc.a, tc.b)
				}
				return
			}
			if err != nil {
				t.Fatalf("Div(%d, %d): unexpected error: %v", tc.a, tc.b, err)
			}
			if got != tc.want {
				t.Errorf("Div(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
