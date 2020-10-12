package sublist

// Relation represents the possible relations that can exist between
// 2 compared lists (sublist, superlist, equal, unequal).
type Relation string

// Sublist takes two integer slices a and b and returns the Relation
// between the two slices (sublist, superlist, equal, unequal).
func Sublist(a, b []int) Relation {
	if len(a) < len(b) {
		if len(a) == 0 {
			return "sublist"
		}
		for i := 0; i+len(a) <= len(b); i++ {
			if equal(a, b[i:i+len(a)]) {
				return "sublist"
			}
		}
		return "unequal"
	}

	if len(a) > len(b) {
		if len(b) == 0 {
			return "superlist"
		}
		for i := 0; i+len(b) <= len(a); i++ {
			if equal(b, a[i:i+len(b)]) {
				return "superlist"
			}
		}
		return "unequal"
	}

	if equal(a, b) {
		return "equal"
	}

	return "unequal"
}

func equal(a, b []int) bool {
	if len(a) == len(b) {
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	return false
}
