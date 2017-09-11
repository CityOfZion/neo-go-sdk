package neo_test

import (
	"testing"

	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/stretchr/testify/assert"
)

func BenchmarkWIF(b *testing.B) {
	in := "L1QqQJnpBwbsPGAuutuzPTac8piqvbR1HRjrY5qHup48TBCBFe4g"

	for i := 0; i < b.N; i++ {
		wif := neo.NewWIF(in)
		wif.ToPrivateKey()
	}
}

func TestWIF(t *testing.T) {
	t.Run(".ToPrivateKey()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				in          string
				out         string
			}{
				{
					description: "Valid",
					in:          "L1QqQJnpBwbsPGAuutuzPTac8piqvbR1HRjrY5qHup48TBCBFe4g",
					out:         "7d128a6d096f0c14c3a25a2b0c41cf79661bfcb4a8cc95aaaea28bde4d732344",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.description, func(t *testing.T) {
					wif := neo.NewWIF(testCase.in)
					privateKey, err := wif.ToPrivateKey()
					assert.Nil(t, err)
					assert.Equal(t, testCase.out, privateKey.Value)
					assert.Equal(t, 64, len(privateKey.Value))
				})
			}
		})
	})
}
