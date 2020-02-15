package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency makes concurrent calls to the Frequency function
// and returns a FreqMap containing letter counts for a given list
// of strings.
func ConcurrentFrequency(inputs []string) FreqMap {
	ch := make(chan FreqMap)
	freq := make(FreqMap)

	for _, text := range inputs {
		go func(s string, c chan FreqMap) {
			c <- Frequency(s)
		}(text, ch)
	}

	for range inputs {
		for k, v := range <-ch {
			freq[k] += v
		}
	}
	return freq
}
