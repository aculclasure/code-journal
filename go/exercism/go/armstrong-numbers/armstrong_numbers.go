package armstrong

// IsNumber accepts a number as an int and returns a bool
// if the number is an Armstrong number.
func IsNumber(n int) bool {
	digits := []int{}
	remainder := n

	for remainder > 0 {
		digits = append(digits, remainder%10)
		remainder /= 10
	}

	armstrongSum := 0
	for _, digit := range digits {
		armstrongSum += pow(digit, len(digits))
	}
	return armstrongSum == n
}

func pow(base, exp int) int {
	result := 1

	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
