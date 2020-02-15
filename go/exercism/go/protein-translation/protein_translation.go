// Package protein contains RNA and Codon decoding functions.
package protein

import (
	"errors"
)

var proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	// ErrStop when we encounter one of the sentinel codons
	ErrStop = errors.New("encountered a stop protein")
	// ErrInvalidBase when we encounter an invalid codon
	ErrInvalidBase = errors.New("invalid base")
)

// FromCodon accepts a codon as as a string and returns the corresponding
// protein if there is a match and returns an error otherwise.
func FromCodon(codon string) (string, error) {
	if protein, ok := proteins[codon]; ok {
		if protein == "STOP" {
			return "", ErrStop
		}
		return protein, nil
	}
	return "", ErrInvalidBase
}

// FromRNA returns a list of proteins given an RNA strand.
func FromRNA(rna string) ([]string, error) {
	var aminoAcids []string

	for i := 0; i < len(rna); i += 3 {
		nextCodon := rna[i : i+3]
		aminoAcid, err := FromCodon(nextCodon)
		if err == ErrStop {
			return aminoAcids, nil
		} else if err == ErrInvalidBase {
			return aminoAcids, err
		}
		aminoAcids = append(aminoAcids, aminoAcid)
	}
	return aminoAcids, nil
}
