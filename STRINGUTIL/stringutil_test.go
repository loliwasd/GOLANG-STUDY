package stringutil

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{"hello", "hello", "olleh"},
		{"привет", "привет", "тевирп"},
		{"empty string", "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Reverse(tc.s)
			if got != tc.want {
				t.Errorf("Reverse(%q) = %s, want %s; Additional info hex values = %x   %x", tc.s, got, tc.want, got, tc.want)
			}
		})
	}
}

func TestCountRunes(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want int
	}{
		{"hello", "hello", 5},
		{"привет", "привет", 6},
		{"привет with emoji", "Привет 😊", 8},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := CountRunes(tc.s)
			if got != tc.want {
				t.Errorf("CountRunes(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
