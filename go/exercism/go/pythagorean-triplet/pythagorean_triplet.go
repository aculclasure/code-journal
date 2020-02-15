package pythagorean

// Triplet represents a Pythagorean triplet
type Triplet [3]int

// Range returns a list of Pythagorean triplets with sides in the range
// min to max inclusive.
func Range(min, max int) []Triplet {
	var (
		triplets      []Triplet
		a, b, c, m, n int
	)

	m = 2
	for c <= max {
		for n = 1; n < m; n++ {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if c < min || c > max {
				break
			}
			if a > b {
				a, b = b, a
			}
			if a > min && b > min {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
		m++
	}
	return triplets
}

// Sum returns a list of Pythagorean triplets where the sum of
// all sides is p.
func Sum(p int) []Triplet {
	var triplets []Triplet

	for a := 1; a <= p/3; a++ {
		for b := a + 1; b <= p/2; b++ {
			c := p - (a + b)
			if a*a+b*b == c*c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}
