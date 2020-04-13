package palindrome

import (
	"errors"
)

// Product represents a multiple and the factors that produced the multiple.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products accepts a minimum and maximum factor boundary and returns the minimum
// and maximum palindromic product within the factor range or returns an error
// otherwise.
func Products(minFactor, maxFactor int) (minProduct, maxProduct Product, err error) {
	if minFactor > maxFactor {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	var (
		palindromesSeen = map[int]*Product{}
		min, max        *Product
	)
	for i := minFactor; i <= maxFactor; i++ {
		for j := i; j <= maxFactor; j++ {
			product := i * j
			if isPalindrome(product) {
				if palindrome, seen := palindromesSeen[product]; seen {
					palindrome.Factorizations = append(
						palindrome.Factorizations, [2]int{i, j})
				} else {
					palindromeProduct := &Product{
						Product:        product,
						Factorizations: [][2]int{[2]int{i, j}},
					}
					palindromesSeen[product] = palindromeProduct
					if min == nil && max == nil {
						min, max = palindromeProduct, palindromeProduct
					} else {
						if product < min.Product {
							min = palindromeProduct
						}
						if product > max.Product {
							max = palindromeProduct
						}
					}
				}
			}
		}
	}

	if len(palindromesSeen) < 1 {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	return *min, *max, nil
}

func isPalindrome(n int) bool {
	divisor := 1

	for n/divisor >= 10 {
		divisor *= 10
	}

	for n != 0 {
		leading := n / divisor
		trailing := n % 10

		if leading != trailing {
			return false
		}

		n = (n % divisor) / 10
		divisor /= 100
	}
	return true
}
