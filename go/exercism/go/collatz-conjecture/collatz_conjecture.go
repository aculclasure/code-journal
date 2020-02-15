// Package collatzconjecture contains a function that applies
// the CollatzConjecture to a given starting point number.
package collatzconjecture

import (
	"fmt"
)

// CollatzConjecture returns the number of iterations it takes to get
// to 1 by applying the Collatz Conjecture to a given start point.
func CollatzConjecture(start int) (int, error) {
	var numSteps int

	if start <= 0 {
		return numSteps, fmt.Errorf("illegal start value given: %d", start)
	}
	for start != 1 {
		switch start%2 == 0 {
		case true:
			start /= 2
		default:
			start = (start * 3) + 1
		}
		numSteps++
	}
	return numSteps, nil
}
