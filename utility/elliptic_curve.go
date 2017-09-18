// Original work completed by @vsergeev: https://github.com/vsergeev/btckeygenie
// Expanded upon here under MIT license.

package utility

import (
	"fmt"
	"math/big"
)

type (
	// EllipticCurve represents the parameters of a short Weierstrass equation elliptic
	// curve.
	EllipticCurve struct {
		A  *big.Int
		B  *big.Int
		P  *big.Int
		G  EllipticCurvePoint
		N  *big.Int
		H  *big.Int
		ma ModularArithmetic
	}
)

// NewEllipticCurve creates a new EllipticCurve struct.
func NewEllipticCurve() EllipticCurve {
	return EllipticCurve{
		ma: NewModularArithmetic(),
	}
}

// Add computes R = P + Q on EllipticCurve ec.
func (e *EllipticCurve) Add(P, Q EllipticCurvePoint) (*EllipticCurvePoint, error) {
	var resultPoint EllipticCurvePoint

	if e.isInfinity(P) && e.isInfinity(Q) {
		resultPoint.X = nil
		resultPoint.Y = nil
	} else if e.isInfinity(P) {
		resultPoint.X = new(big.Int).Set(Q.X)
		resultPoint.Y = new(big.Int).Set(Q.Y)
	} else if e.isInfinity(Q) {
		resultPoint.X = new(big.Int).Set(P.X)
		resultPoint.Y = new(big.Int).Set(P.Y)
	} else if P.X.Cmp(Q.X) == 0 && e.ma.Add(P.Y, Q.Y, e.P).Sign() == 0 {
		resultPoint.X = nil
		resultPoint.Y = nil
	} else if P.X.Cmp(Q.X) == 0 && P.Y.Cmp(Q.Y) == 0 && P.Y.Sign() != 0 {
		num := e.ma.Add(
			e.ma.Mul(
				big.NewInt(3),
				e.ma.Mul(P.X, P.X, e.P),
				e.P,
			),
			e.A,
			e.P,
		)
		den := e.ma.Inverse(e.ma.Mul(big.NewInt(2), P.Y, e.P), e.P)
		lambda := e.ma.Mul(num, den, e.P)

		resultPoint.X = e.ma.Sub(
			e.ma.Mul(lambda, lambda, e.P),
			e.ma.Mul(big.NewInt(2), P.X, e.P),
			e.P,
		)
		resultPoint.Y = e.ma.Sub(
			e.ma.Mul(lambda, e.ma.Sub(P.X, resultPoint.X, e.P), e.P),
			P.Y,
			e.P,
		)
	} else if P.X.Cmp(Q.X) != 0 {
		num := e.ma.Sub(Q.Y, P.Y, e.P)
		den := e.ma.Inverse(e.ma.Sub(Q.X, P.X, e.P), e.P)
		lambda := e.ma.Mul(num, den, e.P)

		resultPoint.X = e.ma.Sub(
			e.ma.Sub(
				e.ma.Mul(lambda, lambda, e.P),
				P.X,
				e.P,
			),
			Q.X,
			e.P,
		)

		resultPoint.Y = e.ma.Sub(
			e.ma.Mul(
				lambda,
				e.ma.Sub(P.X, resultPoint.X, e.P),
				e.P,
			),
			P.Y,
			e.P,
		)
	} else {
		return nil, fmt.Errorf("Unsupported point addition: %v + %v", P.Format(), Q.Format())
	}

	return &resultPoint, nil
}

// IsOnCurve checks if point P is on EllipticCurve ec.
func (e *EllipticCurve) IsOnCurve(point EllipticCurvePoint) bool {
	if e.isInfinity(point) {
		return false
	}

	lhs := e.ma.Mul(point.Y, point.Y, e.P)
	rhs := e.ma.Add(
		e.ma.Add(
			e.ma.Exp(point.X, big.NewInt(3), e.P),
			e.ma.Mul(e.A, point.X, e.P),
			e.P,
		),
		e.B,
		e.P,
	)

	if lhs.Cmp(rhs) == 0 {
		return true
	}

	return false
}

// ScalarBaseMult computes Q = k * G on EllipticCurve ec.
func (e *EllipticCurve) ScalarBaseMult(k *big.Int) (*EllipticCurvePoint, error) {
	return e.ScalarMult(k, e.G)
}

// ScalarMult computes Q = k * P on EllipticCurve ec.
func (e *EllipticCurve) ScalarMult(k *big.Int, P EllipticCurvePoint) (*EllipticCurvePoint, error) {
	var R0 EllipticCurvePoint
	var R1 EllipticCurvePoint

	R0.X = nil
	R0.Y = nil
	R1.X = new(big.Int).Set(P.X)
	R1.Y = new(big.Int).Set(P.Y)

	for i := e.N.BitLen() - 1; i >= 0; i-- {
		if k.Bit(i) == 0 {
			x, err := e.Add(R0, R1)
			if err != nil {
				return nil, err
			}
			R1 = *x

			x, err = e.Add(R0, R0)
			if err != nil {
				return nil, err
			}
			R0 = *x
		} else {
			x, err := e.Add(R0, R1)
			if err != nil {
				return nil, err
			}
			R0 = *x

			x, err = e.Add(R1, R1)
			if err != nil {
				return nil, err
			}
			R1 = *x
		}
	}

	return &R0, nil
}

func (e EllipticCurve) isInfinity(point EllipticCurvePoint) bool {
	if point.X == nil && point.Y == nil {
		return true
	}

	return false
}
