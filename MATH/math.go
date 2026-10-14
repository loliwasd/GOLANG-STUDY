package math

// Add returns an integer sum of two integer numbers
func Add(a, b int) int {
	return a + b
}

// Sub returns an integer substitution of two integer numbers
func Sub(a, b int) int {
	return a - b
}

// Max returns integer maximum value among two integer numbers
func Max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
