package strain

// Ints stores a collection of integers
type Ints []int

// Lists stores a collection of integer slices.
type Lists [][]int

// Strings stores a collection of strings
type Strings []string

// Keep accepts a predicate and returns an Ints object
// containing all integers that satisfy the predicate.
func (i Ints) Keep(predicate func(int) bool) Ints {
	var matches Ints

	for _, val := range i {
		if predicate(val) {
			matches = append(matches, val)
		}
	}

	return matches
}

// Discard accepts a predicate and returns an Ints object
// containing all integers that do not satisfy the predicate.
func (i Ints) Discard(predicate func(int) bool) Ints {
	var discards Ints

	for _, val := range i {
		if !predicate(val) {
			discards = append(discards, val)
		}
	}

	return discards
}

// Keep accepts a predicate and returns a Strings object containing
// all strings that satisfy the predicate.
func (s Strings) Keep(predicate func(string) bool) Strings {
	var matches Strings

	for _, val := range s {
		if predicate(val) {
			matches = append(matches, val)
		}
	}

	return matches
}

// Keep accepts a predicate and returns a Lists object that
// contains all lists of integers which satisfy the predicate.
func (l Lists) Keep(predicate func([]int) bool) Lists {
	var matches Lists

	for _, val := range l {
		if predicate(val) {
			matches = append(matches, val)
		}
	}

	return matches
}
