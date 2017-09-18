package neo

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/CityOfZion/neo-go-sdk/utility"
	"golang.org/x/crypto/ripemd160"
)

type (
	// PrivateKey is a struct which holds the value of a NEO address private key, and has
	// a number of utility functions attached to the struct.
	PrivateKey struct {
		bytes []byte
	}
)

// NewPrivateKeyFromWIF creates a PrivateKey struct using a WIF.
func NewPrivateKeyFromWIF(wif string) (*PrivateKey, error) {
	base58 := utility.NewBase58()

	decodedWIF, err := base58.Decode(wif)
	if err != nil {
		return nil, err
	}

	if len(decodedWIF) != 38 {
		return nil, fmt.Errorf(
			"Expected length of decoded WIF to be 38, got: %d", len(decodedWIF),
		)
	}

	if decodedWIF[0] != 0x80 {
		return nil, fmt.Errorf(
			"Expected first byte of decoded WIF to be '0x80', got: %x", decodedWIF[0],
		)
	}

	if decodedWIF[33] != 0x01 {
		return nil, fmt.Errorf(
			"Expected 34th byte of decoded WIF to be '0x01', got: %x", decodedWIF[33],
		)
	}

	subString := decodedWIF[:len(decodedWIF)-4]

	rawFirstSHA := sha256.Sum256([]byte(subString))
	firstSHA := rawFirstSHA[:]

	rawSecondSHA := sha256.Sum256(firstSHA)
	secondSHA := rawSecondSHA[:]

	firstFourBytes := secondSHA[:4]
	lastFourBytes := decodedWIF[len(decodedWIF)-4 : len(decodedWIF)]

	if !bytes.Equal(firstFourBytes, lastFourBytes) {
		return nil, fmt.Errorf("WIF failed checksum validation")
	}

	return &PrivateKey{
		bytes: decodedWIF[1:33],
	}, nil
}

// PublicAddress derives the public key address that is coupled with the private key, and
// outputs it as a string.
func (p PrivateKey) PublicAddress() (string, error) {
	ellipticCurve := p.createEllipticCurve()
	bytesInt := new(big.Int).SetBytes(p.bytes)

	point, err := ellipticCurve.ScalarBaseMult(bytesInt)
	if err != nil {
		return "", err
	}

	if !ellipticCurve.IsOnCurve(*point) {
		return "", errors.New("failed to derive public key using elliptic curve")
	}

	pointXBytes := point.X.Bytes()
	paddedPointXBytes := append(
		bytes.Repeat(
			[]byte{0x00},
			32-len(pointXBytes),
		),
		pointXBytes...,
	)

	var bytes []byte
	if point.Y.Bit(0) == 0 {
		bytes = append([]byte{0x02}, paddedPointXBytes...)
	} else {
		bytes = append([]byte{0x03}, paddedPointXBytes...)
	}

	bytes = append([]byte{0x21}, bytes...)
	bytes = append(bytes, 0xAC)

	sha256H := sha256.New()
	sha256H.Reset()
	sha256H.Write(bytes)
	hashOne := sha256H.Sum(nil)

	ripemd160H := ripemd160.New()
	ripemd160H.Reset()
	ripemd160H.Write(hashOne)
	hashTwo := ripemd160H.Sum(nil)

	var ver uint8 = 0x17

	hashTwo = append([]byte{ver}, hashTwo...)

	sha256H.Reset()
	sha256H.Write(hashTwo)
	hashThree := sha256H.Sum(nil)

	sha256H.Reset()
	sha256H.Write(hashThree)
	hashFour := sha256H.Sum(nil)

	hashTwo = append(hashTwo, hashFour[0:4]...)

	base58 := utility.NewBase58()
	address := base58.Encode(hashTwo)

	return address, nil
}

// Output converts the 32-byte slice representation of the private key to a string.
func (p PrivateKey) Output() string {
	return hex.EncodeToString(p.bytes)
}

// OutputBase64 converts the 32-byte slice representation of the private key to a base64
// encoded string.
func (p PrivateKey) OutputBase64() string {
	return base64.StdEncoding.EncodeToString(p.bytes)
}

func (p PrivateKey) createEllipticCurve() utility.EllipticCurve {
	var ellipticCurve utility.EllipticCurve

	ellipticCurve.P, _ = new(big.Int).SetString(
		"FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF", 16,
	)
	ellipticCurve.A, _ = new(big.Int).SetString(
		"FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFC", 16,
	)
	ellipticCurve.B, _ = new(big.Int).SetString(
		"5AC635D8AA3A93E7B3EBBD55769886BC651D06B0CC53B0F63BCE3C3E27D2604B", 16,
	)
	ellipticCurve.G.X, _ = new(big.Int).SetString(
		"6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296", 16,
	)
	ellipticCurve.G.Y, _ = new(big.Int).SetString(
		"4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5", 16,
	)
	ellipticCurve.N, _ = new(big.Int).SetString(
		"FFFFFFFF00000000FFFFFFFFFFFFFFFFBCE6FAADA7179E84F3B9CAC2FC632551", 16,
	)
	ellipticCurve.H, _ = new(big.Int).SetString("01", 16)

	return ellipticCurve
}
