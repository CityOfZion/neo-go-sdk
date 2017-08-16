package response

type (
	// String represents the JSON schema of a response from a NEO node, where the expected
	// result is a string.
	String struct {
		ID      int    `json:"id"`
		JSONRPC string `json:"jsonrpc"`
		Result  string `json:"result"`
	}
)
