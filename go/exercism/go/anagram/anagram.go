package anagram

import (
	"strings"
)

func getCharCountMap(s string) map[rune]int {
	charCountMap := make(map[rune]int)

	for _, nextChar := range s {
		charCountMap[nextChar]++
	}
	return charCountMap
}

func charMapsEqual(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k1, v1 := range map1 {
		if v2, ok := map2[k1]; !ok {
			return false
		} else if v1 != v2 {
			return false
		}
	}
	return true
}

// Detect returns a list of strings from the candidates argument
// that are anagrams of the given string argument.
func Detect(s string, candidates []string) []string {
	var matches []string
	s = strings.ToLower(s)
	charMapToMatch := getCharCountMap(s)

	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if lowerCandidate == s {
			continue
		}
		candidateMap := getCharCountMap(lowerCandidate)
		if charMapsEqual(charMapToMatch, candidateMap) {
			matches = append(matches, candidate)
		}
	}
	return matches
}
