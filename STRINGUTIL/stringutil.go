package stringutil

// Reverse reverses runes of a given string, returning reversed string
func Reverse(s string) string {

	reversed := []rune(s)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)
}

// CountRunes returns number of runes in a given string
func CountRunes(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
