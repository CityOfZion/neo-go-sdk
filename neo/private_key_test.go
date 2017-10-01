package neo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/stretchr/testify/assert"
)

func TestPrivateKey(t *testing.T) {
	t.Run("NewPrivateKeyFromWIF()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			for i, account := range testAccounts {
				description := fmt.Sprintf("%d", i)
				t.Run(description, func(t *testing.T) {
					privateKey, err := neo.NewPrivateKeyFromWIF(account.wif)
					assert.IsType(t, &neo.PrivateKey{}, privateKey)
					assert.NoError(t, err)
				})
			}
		})
	})

	t.Run(".Output()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			for i, account := range testAccounts {
				description := fmt.Sprintf("%d", i)
				t.Run(description, func(t *testing.T) {
					privateKey, err := neo.NewPrivateKeyFromWIF(account.wif)
					assert.NoError(t, err)

					output := privateKey.Output()
					assert.Equal(t, account.privateKey, output)
				})
			}
		})
	})

	t.Run(".OutputBase64()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			for i, account := range testAccounts {
				description := fmt.Sprintf("%d", i)
				t.Run(description, func(t *testing.T) {
					privateKey, err := neo.NewPrivateKeyFromWIF(account.wif)
					assert.NoError(t, err)

					output := privateKey.OutputBase64()
					assert.Equal(t, account.privateKeyBase64, output)
				})
			}
		})
	})

	t.Run(".PublicAddress()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			for i, account := range testAccounts {
				description := fmt.Sprintf("%d", i)
				t.Run(description, func(t *testing.T) {
					privateKey, err := neo.NewPrivateKeyFromWIF(account.wif)
					assert.NoError(t, err)

					publicAddress, err := privateKey.PublicAddress()
					assert.NoError(t, err)
					assert.Equal(t, account.publicAddress, publicAddress)
				})
			}
		})
	})

	t.Run(".PublicKey()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			for i, account := range testAccounts {
				description := fmt.Sprintf("%d", i)
				t.Run(description, func(t *testing.T) {
					privateKey, err := neo.NewPrivateKeyFromWIF(account.wif)
					assert.NoError(t, err)

					publicKeyBytes, err := privateKey.PublicKey()
					assert.NoError(t, err)

					publicKey := hex.EncodeToString(publicKeyBytes)

					assert.Equal(t, account.publicKey, publicKey)
				})
			}
		})
	})

	t.Run(".Signature()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			for i, account := range testAccounts {
				description := fmt.Sprintf("%d", i)
				t.Run(description, func(t *testing.T) {
					privateKey, err := neo.NewPrivateKeyFromWIF(account.wif)
					assert.NoError(t, err)

					signatureBytes, err := privateKey.Signature()
					assert.NoError(t, err)

					signature := hex.EncodeToString(signatureBytes)

					assert.Equal(t, account.signature, signature)
				})
			}
		})
	})
}
