package secret

var moves = []struct {
	code uint
	move string
}{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

// Handshake accepts an integer and returns a slice of moves
// that are found within its binary encoding.
func Handshake(val uint) []string {
	var matches []string

	for _, m := range moves {
		if m.code&val == m.code {
			matches = append(matches, m.move)
		}
	}
	if val&16 == 16 {
		reverse(matches)
	}
	return matches
}

func reverse(values []string) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}
