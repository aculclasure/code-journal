package darts

import (
	"math"
)

// Score given x and y cartesian coordinates for a dart throw
// returns the score.
func Score(x, y float64) int {
	distanceFromCenter := math.Sqrt((y * y) + (x * x))

	switch {
	case distanceFromCenter <= 1.0:
		return 10
	case distanceFromCenter > 1.0 && distanceFromCenter <= 5.0:
		return 5
	case distanceFromCenter > 5.0 && distanceFromCenter <= 10.0:
		return 1
	default:
		return 0
	}
}
