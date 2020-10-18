package house

var phrases = []string{
	"the house that Jack built.",
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
}

// Verse returns the nth verse of the song as a string.
func Verse(n int) string {
	return buildVerse(n, n)
}

func buildVerse(n, previousN int) string {
	if n == 1 && n == previousN {
		return "This is " + phrases[n-1]
	}
	if n == 1 {
		return phrases[n-1]
	}
	if n == previousN {
		return "This is " + phrases[n-1] + " " + buildVerse(n-1, n)
	}
	return phrases[n-1] + " " + buildVerse(n-1, n)
}

// Song returns the entire song as a string.
func Song() string {
	song := Verse(1)

	for i := 2; i <= len(phrases); i++ {
		song += "\n\n" + Verse(i)
	}
	return song
}
