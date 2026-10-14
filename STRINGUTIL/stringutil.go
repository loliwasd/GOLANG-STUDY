package stringutil

func Reverse(s string) string {

	_s := []rune(s)

	var new_s []rune = make([]rune, len(_s))

	for i, sr := range _s {
		new_s[len(_s)-i-1] = sr
	}

	return string(new_s)
}

func CountRunes(s string) int {
	var count int

	_s := []rune(s)

	for range _s {
		count += 1
	}

	return count
}
