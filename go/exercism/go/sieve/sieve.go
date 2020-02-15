package sieve

// Sieve determines the prime numbers up to limit n
// using the Sieve of Eratosthenes algorithm and returns
// them as a list of integers.
func Sieve(limit int) []int {
	crossedOut := make([]bool, limit+1)
	var primeNumbers []int

	if limit < 2 {
		return primeNumbers
	}
	for i := 2; i*i <= limit; i++ {
		if crossedOut[i] == false {
			for j := i * 2; j <= limit; j += i {
				crossedOut[j] = true
			}
		}
	}
	for i := 2; i <= limit; i++ {
		if crossedOut[i] == false {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}
