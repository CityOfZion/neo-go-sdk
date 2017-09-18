package utility

import "math/big"

type (
	// ModularArithmetic is a set of helper functions for carrying out specialised
	// mathematical operations in an elliptic curve.
	ModularArithmetic struct{}
)

// NewModularArithmetic creates a new ModularArithmetic struct.
func NewModularArithmetic() ModularArithmetic {
	return ModularArithmetic{}
}

// Add computes: (x + y) % p.
func (m ModularArithmetic) Add(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).Add(x, y)
	z.Mod(z, p)
	return z
}

// Exp computes: (x^e) % p.
func (m ModularArithmetic) Exp(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).Exp(x, y, p)
	return z
}

// Inverse computes: (1/x) % p.
func (m ModularArithmetic) Inverse(x *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).ModInverse(x, p)
	return z
}

// Mul computes: (x * y) % p.
func (m ModularArithmetic) Mul(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	n := new(big.Int).Set(x)
	z = big.NewInt(0)

	for i := 0; i < y.BitLen(); i++ {
		if y.Bit(i) == 1 {
			z = m.Add(z, n, p)
		}
		n = m.Add(n, n, p)
	}

	return z
}

// Sqrt computes: sqrt(x) % p.
func (m ModularArithmetic) Sqrt(x *big.Int, p *big.Int) (z *big.Int) {
	if new(big.Int).Mod(p, big.NewInt(4)).Cmp(big.NewInt(3)) != 0 {
		panic("p is not equal to 3 mod 4!")
	}

	e := new(big.Int).Add(p, big.NewInt(1))
	e = e.Rsh(e, 2)

	z = m.Exp(x, e, p)
	return z
}

// Sub computes: (x - y) % p.
func (m ModularArithmetic) Sub(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).Sub(x, y)
	z.Mod(z, p)
	return z
}
