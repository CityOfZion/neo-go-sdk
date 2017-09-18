package utility

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

type (
	// EllipticCurvePoint represents a point on an elliptic curve.
	EllipticCurvePoint struct {
		X *big.Int
		Y *big.Int
	}
)

// Format encodes the bytes of an EllipticCurvePoint for debugging.
func (p EllipticCurvePoint) Format() string {
	if p.X == nil && p.Y == nil {
		return "(inf,inf)"
	}

	encodedX := hex.EncodeToString(p.X.Bytes())
	encodedY := hex.EncodeToString(p.Y.Bytes())

	return fmt.Sprintf("(%s,%s)", encodedX, encodedY)
}
