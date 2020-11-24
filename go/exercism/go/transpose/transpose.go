package transpose

// Transpose accepts a slice of strings as input and returns their
// transposition as a slice of strings.
func Transpose(input []string) []string {
	transposed := make([]string, 0)

	for i, row := range input {
		for j, char := range row {
			for len(transposed) <= j {
				transposed = append(transposed, "")
			}
			for len(transposed[j]) < i {
				transposed[j] += " "
			}
			transposed[j] += string(char)
		}
	}
	return transposed
}
