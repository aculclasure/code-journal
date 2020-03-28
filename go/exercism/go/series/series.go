package series

// All returns a slice of all substrings of length n found in string s.
func All(n int, s string) []string {
	var results []string

	for i := 0; i+n <= len(s); i++ {
		results = append(results, s[i:i+n])
	}
	return results
}

// UnsafeFirst returns the first substring of length n in string s.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		first = ""
		ok = false
		return
	}

	first = s[:n]
	ok = true
	return
}
