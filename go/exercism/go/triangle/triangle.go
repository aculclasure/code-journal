// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle contains a single function that determines the triangle type
// given 3 sides.
package triangle

import (
	"math"
)

// Kind refers to the type of triangle
type Kind int

const (
	// NaT indicates not a triangle
	NaT = iota
	// Equ indicates an equilateral triangle
	Equ
	// Iso indicates an isosceles triangle
	Iso
	// Sca indicates a scalene triangle
	Sca
)

func isTriangle(a, b, c float64) bool {
	return (!math.IsInf(a, 0) && !math.IsInf(b, 0) && !math.IsInf(c, 0)) &&
		(!math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c)) &&
		a > 0 && b > 0 && c > 0 &&
		(a+b) >= c && (b+c) >= a && (a+c) >= b
}

// KindFromSides accepts 3 side lengths as arguments and returns the type of triangle as a Kind.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if isTriangle(a, b, c) {
		switch {
		case a == b && (a != c && b != c):
			fallthrough
		case a == c && (a != b && c != b):
			fallthrough
		case b == c && (b != a && c != a):
			k = Iso
		case a == b && b == c && a == c:
			k = Equ
		default:
			k = Sca
		}
	} else {
		k = NaT
	}
	return k
}
