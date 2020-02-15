// Package strand contains a function that converts a DNA strand into its RNA complement.
package strand

var rnaComplements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA accepts a string DNA sequence and returns its RNA complement as a string.
func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	for i, dnaNucleotide := range dna {
		rna[i] = rnaComplements[dnaNucleotide]
	}
	return string(rna)
}
