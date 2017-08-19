package neo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/CityOfZion/neo-go-sdk/utility"
)

type (
	// WIF (wallet import format) is a struct that holds a string, which has a number of
	// utility functions on it.
	WIF struct {
		Value string
	}
)

// NewWIF creates a WIF struct.
func NewWIF(s string) WIF {
	return WIF{
		Value: s,
	}
}

// ToPrivateKey converts the WIF to private key format.
func (w WIF) ToPrivateKey() (*PrivateKey, error) {
	base58 := utility.NewBase58()

	decoded, err := base58.Decode(w.Value)
	if err != nil {
		return nil, err
	}

	if len(decoded) != 38 {
		return nil, fmt.Errorf(
			"Expected length of decoded WIF to be 38, got: %d", len(decoded),
		)
	}

	if decoded[0] != 0x80 {
		return nil, fmt.Errorf(
			"Expected first byte of decoded WIF to be '0x80', got: %x", decoded[0],
		)
	}

	if decoded[33] != 0x01 {
		return nil, fmt.Errorf(
			"Expected 34th byte of decoded WIF to be '0x01', got: %x", decoded[33],
		)
	}

	subString := decoded[:len(decoded)-4]

	rawFirstSHA := sha256.Sum256([]byte(subString))
	firstSHA := rawFirstSHA[:]

	rawSecondSHA := sha256.Sum256(firstSHA)
	secondSHA := rawSecondSHA[:]

	firstFourBytes := secondSHA[:4]
	lastFourBytes := decoded[len(decoded)-4 : len(decoded)]

	for i, x := range firstFourBytes {
		if x != lastFourBytes[i] {
			return nil, fmt.Errorf("WIF failed checksum validation")
		}
	}

	privateKey := NewPrivateKey(
		hex.EncodeToString(decoded[1:33]),
	)

	return &privateKey, nil
}
