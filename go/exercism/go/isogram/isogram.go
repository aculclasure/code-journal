// Package isogram contains functions for determining isograms.
package isogram

import (
	"unicode"
)

// IsIsogram returns true if the given word is an isogram
// and returns false otherwise
func IsIsogram(word string) bool {
	letterCounts := make(map[rune]int)
	runes := []rune(word)

	for _, nextRune := range runes {
		if unicode.IsLetter(nextRune) {
			nextRune = unicode.ToUpper(nextRune)
			if _, ok := letterCounts[nextRune]; ok {
				return false
			}
			letterCounts[nextRune] = 1
		}
	}
	return true
}
