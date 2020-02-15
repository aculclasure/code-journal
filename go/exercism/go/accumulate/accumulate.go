package accumulate

// Accumulate accepts a list of strings and a converter function to apply to each
// string in the list and returns a list of converted strings.
func Accumulate(given []string, converter func(string) string) []string {
	converted := make([]string, len(given))
	for i, wordToConvert := range given {
		converted[i] = converter(wordToConvert)
	}
	return converted
}
