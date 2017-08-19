package utility

import (
	"bytes"
	"fmt"
	"math/big"
)

type (
	// Base58 is a encode/decode utility for base58 strings.
	Base58 struct {
		alphabet  [58]byte
		decodeMap map[byte]int64
	}
)

const (
	defaultAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

// NewBase58 creates a new Base58 struct, using the default alphabet.
func NewBase58() Base58 {
	base58 := Base58{}

	copy(base58.alphabet[:], []byte(defaultAlphabet))

	base58.decodeMap = map[byte]int64{}
	for i, b := range []byte(defaultAlphabet) {
		base58.decodeMap[b] = int64(i)
	}

	return base58
}

// Decode decodes the base58 encoded string.
func (b Base58) Decode(s string) ([]byte, error) {
	source := []byte(s)
	startIndex := 0
	buffer := &bytes.Buffer{}

	for i, c := range source {
		if c == b.alphabet[0] {
			if err := buffer.WriteByte(0x00); err != nil {
				return nil, err
			}
		} else {
			startIndex = i
			break
		}
	}

	n := big.NewInt(0)
	div := big.NewInt(58)

	for _, c := range source[startIndex:] {
		charIndex, ok := b.decodeMap[c]
		if !ok {
			return nil, fmt.Errorf(
				"invalid character '%c' when decoding this base58 string: '%s'", c, source,
			)
		}

		n.Add(n.Mul(n, div), big.NewInt(charIndex))
	}

	buffer.Write(n.Bytes())

	return buffer.Bytes(), nil
}
