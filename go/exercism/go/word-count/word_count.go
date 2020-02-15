package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map of words to their respective counts
type Frequency map[string]int

// WordCount accepts a string a returns a count of words
// in the string.
func WordCount(s string) Frequency {
	wordCount := Frequency{}
	re := regexp.MustCompile(`\w+('\w+)?`)

	for _, word := range re.FindAllString(s, -1) {
		wordCount[strings.ToLower(word)]++
	}
	return wordCount
}
