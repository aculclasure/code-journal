package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey accepts a prime number p and returns a private key
// such that 2 <= private-key < p.
func PrivateKey(p *big.Int) *big.Int {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	limit := new(big.Int).Sub(p, big.NewInt(2))
	startKey := new(big.Int)
	return startKey.Rand(seed, limit).Add(startKey, big.NewInt(2))
}

// PublicKey accepts a private key, a prime number modulus, and a prime
// number generator and returns the calculated public key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair accepts a prime number modulus and a prime number generator
// and returns the calculated private key and public key.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey accepts a private key, public key, and prime number modulus
// and returns the calculated secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
