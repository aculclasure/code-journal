package beer

import (
	"fmt"
)

// Verse returns the nth verse of the song an error if the value of n is invalid.
func Verse(n int) (string, error) {
	switch {
	case n < 0 || n > 99:
		return "", fmt.Errorf("got invalid value for n (want 0 <= n <= 99): %d", n)
	case n >= 3:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n",
			n, n, n-1), nil
	case n == 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	case n == 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	default:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
}

// Verses accepts upperBound and lowerBound int arguments specifying the range of song verses
// and returns the verses in that range as a string or returns an error otherwise.
func Verses(upperBound, lowerBound int) (string, error) {
	if upperBound > 99 || lowerBound < 0 || upperBound < lowerBound {
		return "", fmt.Errorf("got invalid upperBound (%d) and/or lowerBound (%d) arguments (want upperBound >= 99, lowerBound <= 0, upperBound > lowerBound)", upperBound, lowerBound)
	}

	verses := ""
	for upperBound >= lowerBound {
		verse, err := Verse(upperBound)
		if err != nil {
			return "", err
		}
		verses += verse + "\n"
		upperBound--
	}
	return verses, nil
}

// Song returns the entire song "99 Bottles of Beer" as a string
func Song() string {
	actual, _ := Verses(99, 0)
	return actual
}
