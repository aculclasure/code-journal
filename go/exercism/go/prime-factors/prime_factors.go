package prime

// Factors computes the prime factors of a given
// number and returns them as a list of int64.
func Factors(number int64) []int64 {
	var i int64
	primeFactors := []int64{}

	if number < 2 {
		return primeFactors
	}

	for i = 2; number/i >= 1; i++ {
		for number%i == 0 {
			number /= i
			primeFactors = append(primeFactors, i)
		}
	}
	return primeFactors
}
