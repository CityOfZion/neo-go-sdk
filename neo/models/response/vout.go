package response

import "github.com/CityOfZion/neo-go-sdk/neo/models"

type (
	// Vout represents the JSON schema of a response from a NEO node, where the expected
	// result is all the data about a transaction output (vout).
	Vout struct {
		ID      int         `json:"id"`
		JSONRPC string      `json:"jsonrpc"`
		Result  models.Vout `json:"result"`
	}
)
