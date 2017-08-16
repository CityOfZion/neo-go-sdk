package response

import "github.com/CityOfZion/neo-go-sdk/neo/models"

type (
	// Block represents the JSON schema of a response from a NEO node, where the expected
	// result is all the data about a particular block.
	Block struct {
		ID      int          `json:"id"`
		JSONRPC string       `json:"jsonrpc"`
		Result  models.Block `json:"result"`
	}
)
