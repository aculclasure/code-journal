package diamond

import (
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < byte('A') || char > byte('Z') {
		return "", fmt.Errorf("got invalid char (want char in range of A-Z): %s", string(char))
	}

	depth := int(char - byte('A'))
	width := (2 * depth) + 1
	diamond := make([]string, width)
	for r := 0; r <= depth; r++ {
		row := make([]rune, width)
		for col := range row {
			ltr := ' '
			if depth+r == col || depth-r == col {
				ltr = rune('A' + r)
			}
			row[col] = ltr
		}
		diamond[r] = string(row)
		diamond[2*depth-r] = string(row)
	}
	return strings.Join(diamond, "\n") + "\n", nil
}
