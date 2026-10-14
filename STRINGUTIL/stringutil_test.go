package stringutil

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{"hello", "olleh"},
		{"привет", "тевирп"},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := Reverse(tc.s)
			if got != tc.want {
				t.Errorf("Reverse(hello) = %s, want %s; Additional info hex values = %x   %x", got, tc.want, got, tc.want)
			}
		})
	}
}

func TestCountRunes(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{"hello", 5},
		{"привет", 6},
		{"Привет 😊", 8},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := CountRunes(tc.s)
			if got != tc.want {
				t.Errorf("CountRunes(hello) = %d, want %d", got, tc.want)
			}
		})
	}
}
