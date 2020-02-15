package summultiples


// SumMultiples computes and returns the sum of unique
// multiples up to but not including the given limit.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	multiplesExist := make(map[int]bool)

	for _, divisor := range divisors {
		if divisor > 0 {
			for m := divisor; m < limit; m += divisor {
				if _, ok := multiplesExist[m]; !ok {
					sum += m
					multiplesExist[m] = true
				}
			}
		}
	}
	return sum
}
