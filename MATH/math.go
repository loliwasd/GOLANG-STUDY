package math

import "errors"

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

// Div returns division of two integer numbers, and an error if second number is zero
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
