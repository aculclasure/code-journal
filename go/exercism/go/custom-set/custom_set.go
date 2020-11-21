package stringset

import (
	"strings"
)

// Set represents a custom set type that holds a collection of unique strings.
type Set map[string]bool

// New returns a new empty Set
func New() Set {
	return make(Set, 0)
}

// NewFromSlice accepts a slice of strings, adds them to a Set, and returns the final Set.
func NewFromSlice(items []string) Set {
	set := make(Set, 0)

	for _, item := range items {
		set[item] = true
	}
	return set
}

// String returns the items of the set as a string.
func (s Set) String() string {
	set := "{"

	for k := range s {
		set += "\"" + k + "\", "
	}
	set = strings.TrimRight(set, ", ") + "}"
	return set
}

// IsEmpty returns true if the Set is empty or returns false otherwise.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns true if the given set contains item or returns false otherwise.
func (s Set) Has(item string) bool {
	if _, ok := s[item]; ok {
		return true
	}
	return false
}

// Subset returns true if s1 is a subset of s2 or returns false otherwise.
func Subset(s1, s2 Set) bool {
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 have no elements in common or returns false otherwise.
func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

// Equal returns true if s1 and s2 have the same elements or returns false otherwise.
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Add adds item to the given set.
func (s Set) Add(item string) {
	s[item] = true
}

// Intersection returns a Set whose items are in both s1 and s2.
func Intersection(s1, s2 Set) Set {
	intersection := New()

	for k := range s1 {
		if s2.Has(k) {
			intersection.Add(k)
		}
	}
	return intersection
}

// Difference returns a set whose items are in s1 but not in s2.
func Difference(s1, s2 Set) Set {
	difference := New()

	for k := range s1 {
		if !s2.Has(k) {
			difference.Add(k)
		}
	}
	return difference
}

// Union returns the union of s1 and s2 as a Set.
func Union(s1, s2 Set) Set {
	var s2Items []string

	for k := range s2 {
		s2Items = append(s2Items, k)
	}
	diff := Difference(s1, s2)
	for k := range diff {
		s2Items = append(s2Items, k)
	}
	return NewFromSlice(s2Items)
}
