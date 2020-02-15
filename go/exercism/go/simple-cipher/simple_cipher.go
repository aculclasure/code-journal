package cipher

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

type caesarCipher uint
type shiftCipher int
type vigCipher struct {
	key string
	i   int
}

const (
	modRange       = 26
	minBoundaryVal = int('a')
)

// NewCaesar returns a Cipher whose Encode and Decode methods use
// the Caesar algorithm.
func NewCaesar() Cipher {
	return caesarCipher(0)
}

func (c caesarCipher) Encode(plainText string) string {
	normalized := strings.ToLower(plainText)
	var ciphered string

	for _, c := range normalized {
		if unicode.IsLetter(c) {
			ciphered += string(getShiftedVal(c, 3))
		}
	}
	return ciphered
}

func (c caesarCipher) Decode(ciphered string) string {
	normalized := strings.ToLower(ciphered)
	var plainText string

	for _, c := range normalized {
		if unicode.IsLetter(c) {
			plainText += string(getShiftedVal(c, -3))
		}
	}
	return plainText
}

// NewShift returns a Cipher whose Encode and Decode
// methods use Shift ciphers with a specified distance.
func NewShift(distance int) Cipher {
	absDistance := int(math.Abs(float64(distance)))

	if absDistance == 0 || absDistance > 25 {
		return nil
	}

	return shiftCipher(distance)
}

func (s shiftCipher) Encode(plainText string) string {
	normalized := strings.ToLower(plainText)
	var ciphered string

	for _, c := range normalized {
		if unicode.IsLetter(c) {
			ciphered += string(getShiftedVal(c, int(s)))
		}
	}
	return ciphered
}

func (s shiftCipher) Decode(ciphered string) string {
	normalized := strings.ToLower(ciphered)
	var plainText string

	for _, c := range normalized {
		if unicode.IsLetter(c) {
			plainText += string(getShiftedVal(c, -1*int(s)))
		}
	}
	return plainText
}

// NewVigenere returns a Cipher whose Encode and Decode
// methods use the Vigenere ciphering algorithm.
func NewVigenere(key string) Cipher {
	re := regexp.MustCompile(`[^a-z]+`)

	if re.MatchString(key) || strings.Count(key, "a") == len(key) {
		return nil
	}

	return vigCipher{
		key: key,
		i:   0,
	}
}

func (v *vigCipher) getNextKeyChar() rune {
	r := []rune(v.key)
	nextChar := r[v.i]

	if v.i+1 == len(v.key) {
		v.i = 0
	} else {
		v.i++
	}
	return nextChar
}

func (v vigCipher) Encode(plainText string) string {
	normalized := strings.ToLower(plainText)
	var ciphered string

	for _, c := range normalized {
		if unicode.IsLetter(c) {
			delta := int((&v).getNextKeyChar()) - minBoundaryVal
			ciphered += string(getShiftedVal(c, delta))
		}
	}

	return ciphered
}

func (v vigCipher) Decode(ciphered string) string {
	normalized := strings.ToLower(ciphered)
	var plainText string

	for _, c := range normalized {
		if unicode.IsLetter(c) {
			delta := -1 * (int((&v).getNextKeyChar()) - minBoundaryVal)
			plainText += string(getShiftedVal(c, delta))
		}
	}
	return plainText
}

func getShiftedVal(c rune, delta int) rune {
	switch {
	case delta >= 0:
		return rune((int(c)+delta-minBoundaryVal)%modRange + minBoundaryVal)
	default:
		return rune(((int(c)+delta)-delta*modRange-minBoundaryVal)%modRange + minBoundaryVal)
	}
}
