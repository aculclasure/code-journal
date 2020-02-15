package grains

import (
	"fmt"
	"math/big"
)

// Square computes and returns the number of sand grains for a
// given square it is within allowable bounds and returns an
// error otherwise
func Square(val int) (uint64, error) {
	if val < 1 || val > 64 {
		return 0, fmt.Errorf("received illegal val %d", val)
	}

	x := big.NewInt(2)
	y := big.NewInt(int64(val - 1))
	mod := big.NewInt(0)
	return x.Exp(x, y, mod).Uint64(), nil
}

// Total computes and returns the sum of powers of 2
// on the chessboard assuming each of the 64 squares
// represents a power of 2
func Total() uint64 {
	x := big.NewInt(2)
	y := big.NewInt(65)
	mod := big.NewInt(0)

	// Sum of n powers of 2 is 2 ^ (n+1) - 1
	return x.Exp(x, y, mod).Uint64() - 1
}
