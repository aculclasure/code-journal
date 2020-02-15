package pangram

import (
	"unicode"
)

const alphabetLength int = 26

// IsPangram returns true if the provided string argument
// is a pangram (uses every letter of the alphabet at least once)
func IsPangram(sentence string) bool {
	letters := make(map[rune]bool)

	for _, letter := range sentence {
		letter = unicode.ToLower(letter)
		if letter >= 'a' && letter <= 'z' {
			letters[letter] = true
		}
	}
	return len(letters) == alphabetLength
}
