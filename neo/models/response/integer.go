package response

type (
	// Integer represents the JSON schema of a response from a NEO node, where the expected
	// result is an integer.
	Integer struct {
		ID      int    `json:"id"`
		JSONRPC string `json:"jsonrpc"`
		Result  int64  `json:"result"`
	}
)
