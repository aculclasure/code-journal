// Package proverb contains functions that generate proverb messages given word inputs.
package proverb

import (
	"fmt"
)

// Proverb generates and returns a list of proverbs given a list of words.
func Proverb(rhyme []string) []string {
	proverbs := make([]string, len(rhyme))

	for i, phrase := range rhyme {
		if i == 0 {
			proverbs[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", phrase)
		}
		if i+1 < len(rhyme) {
			proverbs[i] = fmt.Sprintf("For want of a %s the %s was lost.", phrase, rhyme[i+1])
		}
	}

	return proverbs
}
