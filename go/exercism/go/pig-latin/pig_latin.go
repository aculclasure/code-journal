package piglatin

import (
	"regexp"
	"strings"
)

const (
	vowelPattern       = `^([aeiou]|xr|yt)`
	consonantPattern   = `^([bcdfghjklmnpqrstvwz]+|x[bcdfghjklmnpqstvwyz]*|y[bcdfghjklmnpqrsvwz]*)`
	consonantQuPattern = `^([bcdfghjklmnprstvwz]*|x[bcdfghjklmnpstvwyz]*|y[bcdfghjklmnprsvwz]*)qu`
	consonantYPattern  = `^([bcdfghjklmnpqrstvwz]+|x[bcdfghjklmnpqstvwz]*)y`
)

// Sentence takes a plain text sentence and returns its
// Pig Latin equivalent.
func Sentence(plainText string) string {
	var pl []string
	var consonantSound string

	for _, word := range strings.Fields(plainText) {
		switch {
		case regexp.MustCompile(vowelPattern).MatchString(word):
			pl = append(pl, word+"ay")
		case regexp.MustCompile(consonantQuPattern).MatchString(word):
			consonantSound = getConsonantSound(word, consonantQuPattern) + "qu"
			pl = append(pl, word[len(consonantSound):]+consonantSound+"ay")
		case regexp.MustCompile(consonantYPattern).MatchString(word):
			consonantSound = getConsonantSound(word, consonantYPattern)
			pl = append(pl, word[len(consonantSound):]+consonantSound+"ay")
		case regexp.MustCompile(consonantPattern).MatchString(word):
			consonantSound = getConsonantSound(word, consonantPattern)
			pl = append(pl, word[len(consonantSound):]+consonantSound+"ay")
		}
	}
	return strings.Join(pl, " ")
}

func getConsonantSound(s, pattern string) string {
	//return regexp.MustCompile(pattern).FindString(s)
	return regexp.MustCompile(pattern).FindStringSubmatch(s)[1]
}
