// Package hamming contains a function that computes the hamming distance of DNA sequences
package hamming

import (
	"errors"
)

// Distance calculates and returns the hamming distance between 2 strings representing
// DNA sequences.
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New("func expects string arguments a,b to be equal length")
	}

	hammingDistance := 0
	for i := range ar {
		if ar[i] != br[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
