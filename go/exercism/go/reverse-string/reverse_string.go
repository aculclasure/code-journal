// Package reverse contains a function for reversing strings
package reverse

// Reverse accepts a string and returns a new string
// that is the reverse of the argument.
func Reverse(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
