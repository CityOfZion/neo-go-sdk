package neo_test

import (
	"testing"

	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	nodeURI := selectTestNode()

	t.Run("NewClient()", func(t *testing.T) {
		client := neo.NewClient(nodeURI)

		assert.Equal(t, nodeURI, client.NodeURI)
		assert.IsType(t, neo.Client{}, client)
	})

	t.Run(".GetBestBlockHash()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)
			blockHash, err := client.GetBestBlockHash()

			assert.NoError(t, err)
			assert.NotEmpty(t, blockHash)
		})
	})

	t.Run(".GetBlockByHash()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testBlock := range testBlocks {
				t.Run(testBlock.id, func(t *testing.T) {
					block, err := client.GetBlockByHash(testBlock.hash)

					assert.NoError(t, err)
					assert.Equal(t, testBlock.hash, block.Hash)
					assert.Equal(t, testBlock.index, block.Index)
					assert.Equal(t, testBlock.merkleRoot, block.Merkleroot)
					assert.Len(t, block.Transactions, testBlock.numberOfTransactions)
				})
			}
		})
	})

	t.Run(".GetBlockByIndex()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testBlock := range testBlocks {
				t.Run(testBlock.id, func(t *testing.T) {
					block, err := client.GetBlockByIndex(testBlock.index)

					assert.NoError(t, err)
					assert.Equal(t, testBlock.index, block.Index)
					assert.Equal(t, testBlock.hash, block.Hash)
					assert.Equal(t, testBlock.merkleRoot, block.Merkleroot)
					assert.Len(t, block.Transactions, testBlock.numberOfTransactions)
				})
			}
		})
	})

	t.Run(".GetBlockCount()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)
			blockCount, err := client.GetBlockCount()

			assert.NoError(t, err)
			assert.NotEmpty(t, blockCount)
		})
	})

	t.Run(".GetBlockHash()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testBlockHash := range testBlockHashes {
				t.Run(testBlockHash.id, func(t *testing.T) {
					blockHash, err := client.GetBlockHash(testBlockHash.index)

					assert.NoError(t, err)
					assert.NotEmpty(t, blockHash)
					assert.Equal(t, testBlockHash.hash, blockHash)
				})
			}
		})
	})

	t.Run(".GetConnectionCount()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)
			blockCount, err := client.GetConnectionCount()

			assert.NoError(t, err)
			assert.NotEmpty(t, blockCount)
		})
	})

	t.Run(".GetStorage()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			storage, err := client.GetStorage(
				"0xecc6b20d3ccac1ee9ef109af5a7cdb85706b1df9",
				"totalSupply",
			)

			assert.NoError(t, err)
			assert.Equal(t, "0072ef3e2597e201", storage)
		})
	})

	t.Run(".GetTransaction()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testTransaction := range testTransactions {
				t.Run(testTransaction.id, func(t *testing.T) {
					transaction, err := client.GetTransaction(testTransaction.hash)

					assert.NoError(t, err)
					assert.Equal(t, testTransaction.hash, transaction.ID)
					assert.Equal(t, testTransaction.size, transaction.Size)
				})
			}
		})
	})

	t.Run(".GetTransactionOutput()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testTransactionOutput := range testTransactionOutputs {
				t.Run(testTransactionOutput.id, func(t *testing.T) {
					transactionOutput, err := client.GetTransactionOutput(
						testTransactionOutput.hash,
						testTransactionOutput.index,
					)

					assert.NoError(t, err)
					assert.Equal(t, testTransactionOutput.asset, transactionOutput.Asset)
					assert.Equal(t, testTransactionOutput.value, transactionOutput.Value)
				})
			}
		})
	})

	t.Run(".GetUnconfirmedTransactions()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)
			_, err := client.GetUnconfirmedTransactions()

			assert.NoError(t, err)
		})
	})

	t.Run(".Ping()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)
			ok := client.Ping()

			assert.True(t, ok)
		})

		t.Run("SadCase", func(t *testing.T) {
			for _, testPing := range testPings {
				t.Run(testPing.description, func(t *testing.T) {
					client := neo.NewClient(testPing.uri)
					ok := client.Ping()

					assert.False(t, ok)
				})
			}
		})

		t.Run(".ValidateAddress()", func(t *testing.T) {
			t.Run("HappyCase", func(t *testing.T) {
				client := neo.NewClient(nodeURI)

				for _, testAccount := range testAccounts {
					t.Run(testAccount.publicAddress, func(t *testing.T) {
						isValid, err := client.ValidateAddress(testAccount.publicAddress)

						assert.NoError(t, err)
						assert.True(t, isValid)
					})
				}
			})

			t.Run("SadCase", func(t *testing.T) {
				client := neo.NewClient(nodeURI)
				isValid, err := client.ValidateAddress("wake-up-neo")

				assert.NoError(t, err)
				assert.False(t, isValid)
			})
		})
	})
}
